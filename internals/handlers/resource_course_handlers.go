package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/n0o01lh/llp/internals/core/domain"
	"github.com/n0o01lh/llp/internals/core/ports"
)

type ResourceCourseHandlers struct {
	resourceCourseService ports.ResourceCourseService
}

func NewResourceCourseHandlers(rcService ports.ResourceCourseService) *ResourceCourseHandlers {
	return &ResourceCourseHandlers{
		resourceCourseService: rcService,
	}
}

var _ ports.ResourceCourseHandlers = (*ResourceCourseHandlers)(nil)

func (h *ResourceCourseHandlers) AddResourceToCourse(ctx *fiber.Ctx) error {

	var requestBody map[string]interface{}
	var serviceResponse *domain.ResourceCourseResponse
	error := json.Unmarshal(ctx.Body(), &requestBody)

	if error != nil {
		log.Error("Error parsing request body ", error)
		return error
	}

	log.Info("request body ", requestBody)
	// get the resource id
	resourceId := requestBody["resource_id"].(float64)
	// get the course id
	courseId := requestBody["course_id"].(float64)

	result, error := h.resourceCourseService.AddResourceToCourse(uint(resourceId), uint(courseId))

	if error != nil {
		log.Error(error)
		//return error
	}

	serviceResponse = &domain.ResourceCourseResponse{ResourceCourse: result, Error: fmt.Sprintf("%v", error)}

	ctx.JSON(serviceResponse)
	ctx.Response().SetStatusCode(http.StatusOK)

	return nil
}

func (h *ResourceCourseHandlers) AsignCourseToResources(ctx *fiber.Ctx) error {
	var requestBody map[string]any
	error := json.Unmarshal(ctx.Body(), &requestBody)

	if error != nil {
		log.Error("Error parsing request body: ", error)
		return error
	}

	resources, _ := requestBody["resources"].([]any)
	course_id := requestBody["course_id"].(float64)

	courseUpdated, error := h.resourceCourseService.AsignCourseToResources(resources, uint(course_id))

	if error != nil {
		log.Error(error)
		return error
	}

	ctx.JSON(courseUpdated)
	ctx.Response().SetStatusCode(http.StatusOK)

	return nil
}

func (h *ResourceCourseHandlers) RemoveResourceFromCourse(ctx *fiber.Ctx) error {
	var requestBody map[string]any
	error := json.Unmarshal(ctx.Body(), &requestBody)

	if error != nil {
		log.Error("Error parsing request body: ", error)
		return error
	}

	resourceId := requestBody["resource_id"].(float64)
	courseId := requestBody["course_id"].(float64)

	log.Debug("Course Id: ", courseId, " Resource Id: ", resourceId)

	error = h.resourceCourseService.RemoveResourceFromCourse(uint(resourceId), uint(courseId))

	if error != nil {
		log.Error("Error deleting resource from course: ", error)
		return error
	}

	ctx.Response().SetStatusCode(http.StatusOK)

	return nil
}
