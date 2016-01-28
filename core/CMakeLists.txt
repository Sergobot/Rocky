project(RockyCore)

INCLUDE_DIRECTORIES($(CMAKE_CURRENT_SOURCE_DIR))

set(RockyCore_SRCS
    Game.cpp
)

add_library(rockycore SHARED ${RockyCore_SRCS})

install(FILES
    Game.h
    DESTINATION ${CMAKE_INSTALL_PREFIX}/include/Rocky
)