package data

import (
	"AIChat/internal/biz"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
)

type aichatRepo struct {
	data *Data
}

func NewAIChatRepo(data *Data) biz.AIChatRepo {
	return &aichatRepo{
		data: data,
	}
}

func (r *aichatRepo) CloudflareStreamGetAIChat(ctx context.Context, messages *biz.AIChatRequest, ch chan *biz.AIChatRespond, modelKind string) error {

	// Checking whether the modelKind exists.
	modelAddr := ""
	if v, ok := r.data.aiModelTextOnly[modelKind]; !ok {
		r.data.log.Errorf("Model '%s' does not exist.\n", modelKind)
		return errors.New(400, "", "Model does not exist.")
	} else {
		modelAddr = v
	}
	// POST Body serialization
	body, err := json.Marshal(messages)
	if err != nil {
		r.data.log.Error(err)
		return err
	}

	// POST Request settings
	req, err := http.NewRequest("POST", r.data.cfAPI.apiBaseUrl+modelAddr, bytes.NewBuffer(body))
	if err != nil {
		r.data.log.Error(err)
		return err
	}
	req.Header.Add("Authorization", "Bearer "+r.data.cfAPI.token)
	req.Header.Add("Content-Type", "application/json")

	// send request
	resp, err := r.data.http_client.Do(req)
	if err != nil {
		r.data.log.Error(err)
		return err
	}

	// read response
	scanner := bufio.NewScanner(resp.Body)
	defer resp.Body.Close()

	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) > 6 {
			tmp := &biz.AIChatRespond{}
			err := json.Unmarshal(line[6:], tmp) // the respond is text which begins with 'data: ', we have to trim off this part string.
			// r.data.log.Info(tmp)
			if err != nil {
				if text := scanner.Text(); strings.Compare(text, "data: [DONE]") != 0 { // Before the respond finish, it will send the last msg: "data: [DONE]"
					r.data.log.Error(err)
					return err
				}
			}
			ch <- tmp
		}
	}
	return nil
}

func (r *aichatRepo) CloudflareGetAIPaintImg(messages *biz.AIPaintingRequest, modelKind string) ([]byte, error) {

	// Checking whether the modelKind exists.
	modelAddr := ""
	if v, ok := r.data.aiModelTextToImg[modelKind]; !ok {
		r.data.log.Errorf("Model '%s' does not exist.\n", modelKind)
		return nil, errors.New(400, "", "Model does not exist.")
	} else {
		modelAddr = v
	}

	// POST Body serialization
	body, err := json.Marshal(messages)
	if err != nil {
		r.data.log.Error(err)
		return nil, err
	}

	// POST Request settings
	req, err := http.NewRequest("POST", r.data.cfAPI.apiBaseUrl+modelAddr, bytes.NewBuffer(body))
	if err != nil {
		r.data.log.Error(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+r.data.cfAPI.token)
	req.Header.Add("Content-Type", "application/json")

	// send request
	resp, err := r.data.http_client.Do(req)
	if err != nil {
		r.data.log.Error(err)
		return nil, err
	}

	// read response
	// r.data.log.Infof("content-length: %v\n", resp.ContentLength)

	var b = make([]byte, 0, 4096)
	buf := bytes.NewBuffer(b)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}
	// scanner := bufio.NewScanner(resp.Body)
	// size := 0
	// defer resp.Body.Close()
	// for scanner.Scan() {
	// 	v, _ := buf.Write(scanner.Bytes())
	// 	size += v
	// }

	// r.data.log.Infof("len: %v; size: %v\n", buf.Len(), size)
	return buf.Bytes(), nil
}
