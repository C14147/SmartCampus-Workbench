package services

import (
    "fmt"
    "time"
    "smartcampus/internal/repositories"
    "smartcampus/pkg/cache"
)

type PaginationResponse struct {
    Data  interface{} `json:"data"`
    Total int         `json:"total"`
    Page  int         `json:"page"`
}

type CourseService struct {
    CourseRepo repositories.CourseRepository
    Cache      cache.Cache
}

func (s *CourseService) GetTeacherCourses(teacherID string, page, pageSize int) (*PaginationResponse, error) {
    cacheKey := fmt.Sprintf("teacher_courses:%s:%d:%d", teacherID, page, pageSize)

    var result PaginationResponse
    if err := s.Cache.Get(cacheKey, &result); err == nil {
        return &result, nil
    }

    courses, total, err := s.CourseRepo.FindByTeacherID(teacherID, page, pageSize)
    if err != nil {
        return nil, err
    }

    result = PaginationResponse{
        Data:  courses,
        Total: total,
        Page:  page,
    }

    s.Cache.Set(cacheKey, result, 5*time.Minute)

    return &result, nil
}
