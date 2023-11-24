package metadata

import (
	"bytes"
	"fmt"
	"net/http"
	"os"
)

func (m *MetadataService) uploadHandler(res http.ResponseWriter, req *http.Request) {
	fmt.Println("req: ", req)
	file, handler, err := req.FormFile("file")

	if err != nil {
		fmt.Println("error creating formfile: ", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer file.Close()

	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("error opening file: ", err)
		res.WriteHeader(http.StatusInternalServerError)
		return
	}

	defer func() {
		err = os.Remove(handler.Filename)
		if err != nil {
			fmt.Println("error opening file: ", err)
			res.WriteHeader(http.StatusInternalServerError)
		}
		f.Close()
	}()

	buf := bytes.NewBuffer(nil)
}
