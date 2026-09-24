package dto

type PageRequest struct {
    Page  int
    Limit int
}

type PageMeta struct {
    Page       int   `json:"page"`
    Limit      int   `json:"limit"`
    Total      int64 `json:"total"`
    TotalPages int   `json:"totalPages"`
}