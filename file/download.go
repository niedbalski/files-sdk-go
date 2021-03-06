package file

import (
	"fmt"
	"os"
	"path/filepath"

	files_sdk "github.com/Files-com/files-sdk-go"
	"github.com/Files-com/files-sdk-go/folder"
	"github.com/Files-com/files-sdk-go/lib"
	"github.com/zenthangplus/goccm"
)

func (c *Client) DownloadToFile(params files_sdk.FileDownloadParams, filePath string) (files_sdk.File, error) {
	out, err := os.Create(filePath)
	if err != nil {
		return files_sdk.File{}, err
	}
	params.Writer = out
	return c.Download(params)
}

func DownloadToFile(params files_sdk.FileDownloadParams, filePath string) (files_sdk.File, error) {
	return (&Client{}).DownloadToFile(params, filePath)
}

type fileDownload struct {
	files_sdk.File
	destination string
	error
}

func (c *Client) DownloadFolder(params files_sdk.FolderListForParams, rootDestination string, reporters ...func(incDownloadedBytes int64, file files_sdk.File, destination string, err error)) error {
	goc := goccm.New(c.Config.MaxConcurrentConnections())
	files := c.index(params)
	signal := make(chan bool)
	for _, entity := range files {
		go func(entity Entity) {
			goc.Wait()
			file := files_sdk.File{Path: entity.file.Path, Size: entity.file.Size, Type: entity.file.Type}
			destinationPath := filepath.Join(rootDestination, file.Path)
			dir, _ := filepath.Split(destinationPath)
			_, err := os.Stat(dir)
			if os.IsNotExist(err) {
				os.MkdirAll(dir, 0755)
			}
			out, err := os.Create(destinationPath)
			if err != nil {
				if len(reporters) > 0 {
					reporters[0](0, file, destinationPath, err)
				}
			}
			params := files_sdk.FileDownloadParams{Path: file.Path}
			writer := lib.ProgressWriter{Writer: out}
			writer.ProgressWatcher = func(incDownloadedBytes int64) {
				if len(reporters) > 0 {
					reporters[0](incDownloadedBytes, file, destinationPath, entity.error)
				}
			}
			params.Writer = writer
			writer.ProgressWatcher(0)
			newFile, err := c.Download(params)
			if len(reporters) > 0 && err != nil {
				reporters[0](0, newFile, destinationPath, err)
			}
			signal <- true
			goc.Done()
		}(entity)

	}
	for range files {
		<-signal
	}
	return nil
}

func (c *Client) downloadNode(params files_sdk.FolderListForParams, rootDestination string, queuedDownloads *int, goc goccm.ConcurrencyManager, reporter *func(incDownloadedBytes int64, file files_sdk.File, destination string, err error), fileChannel *chan fileDownload) error {
	folderClient := folder.Client{Config: c.Config}
	it, err := folderClient.ListFor(params)

	if err != nil {
		return err
	}

	action := func(entry files_sdk.Folder) {
		destinationPath := filepath.Join(rootDestination, entry.Path)
		download := fileDownload{destination: destinationPath, File: files_sdk.File{Path: entry.Path, Type: entry.Type, Size: entry.Size}}
		switch entry.Type {
		case "directory":
			_, err := os.Stat(destinationPath)
			if os.IsNotExist(err) {
				os.MkdirAll(destinationPath, 0755)
			}
			err = c.downloadNode(files_sdk.FolderListForParams{Path: entry.Path}, rootDestination, queuedDownloads, goc, reporter, fileChannel)
			if err != nil {
				download.error = err
			}
		case "file":
			dir, _ := filepath.Split(destinationPath)
			_, err := os.Stat(dir)
			if os.IsNotExist(err) {
				os.MkdirAll(dir, 0755)
			}
			out, err := os.Create(destinationPath)
			if err != nil {
				download.error = err
				return
			}
			params := files_sdk.FileDownloadParams{Path: entry.Path}
			writer := lib.ProgressWriter{Writer: out}
			writer.ProgressWatcher = func(incDownloadedBytes int64) {
				(*reporter)(incDownloadedBytes, download.File, download.destination, download.error)
			}
			params.Writer = writer
			*queuedDownloads++
			go func() {
				goc.Wait()
				writer.ProgressWatcher(0)

				file, err := c.Download(params)
				completedDownload := fileDownload{File: file, error: err}
				*fileChannel <- completedDownload
				goc.Done()
			}()
		default:
			if it.Err() != nil {
				download.error = it.Err()
			} else {
				download.error = fmt.Errorf("unknown file type %v", entry.Type)
			}
		}
	}

	for it.Next() {
		entry := it.Folder()
		action(entry)
	}
	return nil
}

type Entity struct {
	file files_sdk.Folder
	error
}

func (c *Client) index(params files_sdk.FolderListForParams) []Entity {
	var files []Entity
	folderClient := folder.Client{Config: c.Config}
	it, err := folderClient.ListFor(params)

	if err != nil {
		files = append(files, Entity{file: files_sdk.Folder{Path: params.Path, Type: "error"}, error: err})
	}

	for it.Next() {
		entry := it.Folder()
		switch entry.Type {
		case "directory":
			files = append(files, c.index(files_sdk.FolderListForParams{Path: entry.Path})...)
		case "file":
			files = append(files, Entity{file: entry})
		default:
			if it.Err() != nil {
				files = append(files, Entity{file: entry, error: it.Err()})
			} else {
				files = append(files, Entity{file: entry, error: fmt.Errorf("unknown file type %v", entry.Type)})
			}
		}
	}
	return files
}
