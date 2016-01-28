# Rocky version
set(ROCKY_VERSION_MAJOR 0)
set(ROCKY_VERSION_MINOR 0)
set(ROCKY_VERSION_PATCH 1)
set(ROCKY_VERSION ${ROCKY_VERSION_MAJOR}.${ROCKY_VERSION_MINOR}.${ROCKY_VERSION_PATCH})

cmake_minimum_required(VERSION 2.6)
project(Rocky)

set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -fno-elide-constructors -pedantic-errors -ansi -Werror -Wextra -Winit-self")
set(CMAKE_CXX_FLAGS "${CMAKE_CXX_FLAGS} -Woverloaded-virtual -Wuninitialized -Wmissing-declarations -Winit-self -std=c++11")

add_subdirectory(core)
