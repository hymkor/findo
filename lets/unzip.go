package main

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"os"
	"path"
)

func deepMkdir(folder string) error {
	finfo, err := os.Stat(folder)
	if err == nil {
		if finfo.IsDir() {
			return nil
		} else {
			return fmt.Errorf("%s: Not Directory", folder)
		}
	} else {
		parent := path.Dir(folder)
		if err := deepMkdir(parent); err != nil {
			return err
		}
		if _, err2 := os.Stat(folder); err2 != nil {
			return os.Mkdir(folder, 0666)
		} else {
			return nil
		}
	}
}

func unzip(args_ []string) error {
	listFlag := false
	directory := ""

	args := make([]string, 0, len(args_))
	for i := 0; i < len(args_); i++ {
		switch args_[i] {
		case "-l":
			listFlag = true
		case "-d":
			if i+1 >= len(args_) {
				return errors.New("-d: requires exnapd directory")
			}
			i++
			directory = args_[i]
		default:
			args = append(args, args_[i])
		}
	}
	if len(args) < 1 {
		return errors.New("unzip requires zipfilename")
	}
	zipFileName := args[0]
	reader, readerErr := os.Open(zipFileName)
	if readerErr != nil {
		return readerErr
	}
	defer reader.Close()
	finfo, finfoErr := reader.Stat()
	if finfoErr != nil {
		return finfoErr
	}
	zipReader, zipReaderErr := zip.NewReader(reader, finfo.Size())
	if zipReaderErr != nil {
		return zipReaderErr
	}
	if len(directory) > 0 {
		os.Chdir(directory)
	}
	files := map[string]bool{}
	for _, fname := range args[1:] {
		files[fname] = true
	}
	for _, f := range zipReader.File {
		if len(files) > 0 && !files[f.Name] {
			continue
		}
		if listFlag {
			fmt.Println(f.Name)
			continue
		}
		if f.FileInfo().IsDir() {
			if err := deepMkdir(f.Name); err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s\n", f.Name, err.Error())
			}
			fmt.Fprintln(os.Stdout, f.Name)
			continue
		}
		zipFileReader, zipFileReaderErr := f.Open()
		if zipFileReaderErr != nil {
			fmt.Fprintf(os.Stderr, "%s: %s: %s\n",
				zipFileName,
				f.Name,
				zipFileReaderErr.Error(),
			)
		} else {
			if err := deepMkdir(path.Dir(f.Name)); err != nil {
				fmt.Fprintf(os.Stderr, "%s: %s\n", f.Name, err.Error())
			} else if unzipWriter, unzipWriterErr := os.Create(f.Name); unzipWriterErr != nil {
				fmt.Fprintf(os.Stderr, "%s: %s: %s\n",
					zipFileName,
					f.Name,
					unzipWriterErr.Error())
			} else {
				_, err := io.Copy(unzipWriter, zipFileReader)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s: %s: %s\n",
						zipFileName,
						f.Name,
						err.Error())
				} else {
					fmt.Println(f.Name)
				}
				unzipWriter.Close()
			}
		}
		zipFileReader.Close()
	}
	return nil
}
