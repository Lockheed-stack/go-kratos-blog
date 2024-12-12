package data

import (
	"AI_Service/internal/biz"
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"strings"
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

	// POST Body serialization
	body, err := json.Marshal(messages)
	if err != nil {
		r.data.log.Error(err)
		return err
	}
	// POST Request settings
	req, err := http.NewRequest("POST", r.data.cfAPIBaseUrl+model, bytes.NewBuffer(body))
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
