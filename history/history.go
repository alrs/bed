package history

import "github.com/itchyny/bed/buffer"

type History struct {
	entries []*historyEntry
	index   int
}

type historyEntry struct {
	buffer *buffer.Buffer
	offset int64
	cursor int64
}

func NewHistory() *History {
	return &History{index: -1}
}

func (h *History) Push(buffer *buffer.Buffer, offset int64, cursor int64) {
	newEntry := &historyEntry{buffer.Clone(), offset, cursor}
	if len(h.entries)-1 > h.index {
		h.index++
		h.entries[h.index] = newEntry
		h.entries = h.entries[:h.index+1]
	} else {
		h.entries = append(h.entries, newEntry)
		h.index++
	}
}

func (h *History) Undo() (*buffer.Buffer, int, int64, int64) {
	if h.index < 0 {
		return nil, h.index, 0, 0
	}
	if h.index > 0 {
		h.index--
	}
	e := h.entries[h.index]
	return e.buffer.Clone(), h.index, e.offset, e.cursor
}

func (h *History) Redo() (*buffer.Buffer, int64, int64) {
	if h.index == len(h.entries)-1 || h.index < 0 {
		return nil, 0, 0
	}
	h.index++
	e := h.entries[h.index]
	return e.buffer.Clone(), e.offset, e.cursor
}