package data

import (
	"AI_Service/internal/biz"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"

	"github.com/go-kratos/kratos/v2/errors"
)

type aiCloudflareRepo struct {
	data *Data
}

func NewAICloudflareRepo(data *Data) biz.AICloudflareRepo {
	return &aiCloudflareRepo{
		data: data,
	}
}

// StreamTextGeneration implements biz.AICloudflareRepo.
func (r *aiCloudflareRepo) StreamTextGeneration(ctx context.Context, messages *biz.CloudflareAITextGenerationRequest, ch chan *biz.CloudflareAITextGenerationReply, model string) error {

	// Checking whether the modelKind exists.
	modelAddr := ""
	if v, ok := r.data.aiModelTextOnly[model]; !ok {
		r.data.log.Errorf("Model '%s' does not exist.\n", model)
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
	req, err := http.NewRequest("POST", r.data.cfAPIBaseUrl+modelAddr, bytes.NewBuffer(body))
	if err != nil {
		r.data.log.Error(err)
		return err
	}
	req.Header.Add("Authorization", "Bearer "+r.data.cfToken)
	req.Header.Add("Content-Type", "application/json")
	// send request
	resp, err := r.data.cf_httpClient.Do(req)
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
			tmp := &biz.CloudflareAITextGenerationReply{}
			err := json.Unmarshal(line[6:], tmp) // the respond is text which begins with 'data: ', we have to trim off this part string.
			// r.data.log.Info(tmp)
			if err != nil {
				// Before the respond finish, it will send the last msg: "data: [DONE]"
				if text := scanner.Text(); strings.Compare(text, "data: [DONE]") != 0 {
					r.data.log.Error(err)
					return err
				}
			}
			ch <- tmp
		}
	}
	return nil
}

// ImageGeneration implements biz.AICloudflareRepo.
func (r *aiCloudflareRepo) ImageGeneration(messages *biz.CloudflareAIPaintingRequest, model string) ([]byte, error) {

	// Checking whether the modelKind exists.
	modelAddr := ""
	if v, ok := r.data.aiModelTextToImg[model]; !ok {
		r.data.log.Errorf("Model '%s' does not exist.\n", model)
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
	req, err := http.NewRequest("POST", r.data.cfAPIBaseUrl+modelAddr, bytes.NewBuffer(body))
	if err != nil {
		r.data.log.Error(err)
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+r.data.cfToken)
	req.Header.Add("Content-Type", "application/json")

	// send request
	resp, err := r.data.cf_httpClient.Do(req)
	if err != nil {
		r.data.log.Error(err)
		return nil, err
	}
	// read response
	// r.data.log.Infof("content-length: %v\n", resp.ContentLength)

	var b = make([]byte, 0, 1024*500) // 500KB
	buf := bytes.NewBuffer(b)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
