# \CourseApi

All URIs are relative to *http://api.itu.dk/v2*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CourseCourseIdDelete**](CourseApi.md#CourseCourseIdDelete) | **Delete** /course/{courseId} | 
[**CourseCourseIdGet**](CourseApi.md#CourseCourseIdGet) | **Get** /course/{courseId} | 
[**CourseCourseIdPut**](CourseApi.md#CourseCourseIdPut) | **Put** /course/{courseId} | 
[**CoursePost**](CourseApi.md#CoursePost) | **Post** /course | 


# **CourseCourseIdDelete**
> Course CourseCourseIdDelete(ctx, courseId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **courseId** | **int32**| The ID of the course | 

### Return type

[**Course**](Course.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CourseCourseIdGet**
> Course CourseCourseIdGet(ctx, courseId)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **courseId** | **int32**| The ID of the course | 

### Return type

[**Course**](Course.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CourseCourseIdPut**
> Course CourseCourseIdPut(ctx, courseId, updatedFields)


### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **courseId** | **int32**| The ID of the course | 
  **updatedFields** | [**UpdatedFields**](.md)| The fields to update. You only need to add the fields that need updating. | 

### Return type

[**Course**](Course.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **CoursePost**
> Course CoursePost(ctx, newCourse)


Create a new course

### Required Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
  **newCourse** | [**NewCourse**](.md)| The new course | 

### Return type

[**Course**](Course.md)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: Not defined
 - **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

