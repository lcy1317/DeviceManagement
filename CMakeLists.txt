cmake_minimum_required(VERSION 3.7)

SET(CMAKE_BUILD_TYPE_INIT Release)

project(RGsystem LANGUAGES NONE)

find_program(GO_EXECUTABLE go)
if (NOT GO_EXECUTABLE)
  message(FATAL_ERROR "Cannot find the go")
endif()

find_program(SWAGGER_EXECUTABLE swagger)
if (NOT SWAGGER_EXECUTABLE)
  message(FATAL_ERROR "Cannot find the swagger")
endif()

enable_testing()

add_subdirectory(golang)