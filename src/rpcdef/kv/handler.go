package kv

import "sync"

type Handler struct {
	mu sync.Mutex
	kv map[string]string
}

func (h *Handler) Put(request PutRequest, reply *PutReply) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.kv[request.Key] = request.Value
	reply.Err = OK
	return nil
}

func (h *Handler) Get(request GetRequest, reply *GetReply) error {
	h.mu.Lock()
	defer h.mu.Unlock()
	value, ok := h.kv[request.Key]
	reply.Value = value
	if !ok {
		reply.Err = NotFound
	} else {
		reply.Err = OK
	}
	return nil
}
