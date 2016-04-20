// Copyright (c) Sergey Popov <sergobot@protonmail.com>
// See LICENSE for details.

package rocky

import (
	"fmt"
	"strings"

	"github.com/go-gl/gl/v3.3-core/gl"
)

// ShaderType variables store type of a shader: vertex or fragment
type ShaderType int

const (
	VertexShader   ShaderType = iota
	FragmentShader ShaderType = iota
)

// Shader holds following information about a shader:
// - Its type
// - OpenGL shader ID
// - Shader is compiled or not.
type Shader struct {
	// Type of a shader: vertex or fragment
	t ShaderType

	// OpenGL shader ID
	shader uint32

	// Is true only if the shader is compiled
	compiled bool
}

// Compile compiles a shader from its sources
func (s *Shader) Compile(source string, t ShaderType) error {
	s.compiled = false
	if t == VertexShader {
		s.shader = gl.CreateShader(gl.VERTEX_SHADER)
	} else if t == FragmentShader {
		s.shader = gl.CreateShader(gl.FRAGMENT_SHADER)
	} else {
		return fmt.Errorf("Prevented compiling unsupported shader type: %v", t)
	}
	// If shader type is supported, continue
	s.t = t

	// Convert shader's source code to a C string
	csources, free := gl.Strs(source)
	gl.ShaderSource(s.shader, 1, csources, nil)
	free()
	gl.CompileShader(s.shader)

	// Print compile errors if any
	var status int32
	gl.GetShaderiv(s.shader, gl.COMPILE_STATUS, &status)
	if status == gl.FALSE {
		// Get length of error message
		var logLength int32
		gl.GetShaderiv(s.shader, gl.INFO_LOG_LENGTH, &logLength)

		// Fill compilation log with '\x00's. That's default in C strings
		compLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetShaderInfoLog(s.shader, logLength, nil, gl.Str(compLog))

		return fmt.Errorf("Failed to compile shader: %v", compLog)
	}
	s.compiled = true

	return nil
}

// ShaderProgram contains OpenGL shader program ID and its state:
// is the shader program successfully linked or not.
type ShaderProgram struct {
	// OpenGL shader program ID
	program uint32

	// Is true only if shader program is linked
	linked bool
}

// Link links a shader program from 2 ALREADY COMPILED shaders
func (sp *ShaderProgram) Link(vertex, fragment Shader) error {
	if sp.linked {
		gl.DeleteProgram(sp.program)
		sp.linked = false
	}

	// Shader has wrong type
	if vertex.t != VertexShader {
		return fmt.Errorf("Shader given as a vertex isn't a vertex one")
	}
	if fragment.t != FragmentShader {
		return fmt.Errorf("Shader given as a fragment isn't a fragment one")
	}
	// Shader isn't compiled yet
	if !vertex.compiled {
		return fmt.Errorf("Vertex shader isn't compiled")
	}
	if !fragment.compiled {
		return fmt.Errorf("Fragment shader isn't compiled")
	}

	sp.program = gl.CreateProgram()

	gl.AttachShader(sp.program, vertex.shader)
	gl.AttachShader(sp.program, fragment.shader)
	gl.LinkProgram(sp.program)

	// Print linking errors if any
	var status int32
	gl.GetProgramiv(sp.program, gl.LINK_STATUS, &status)
	if status == gl.FALSE {
		// Get length of error message
		var logLength int32
		gl.GetProgramiv(sp.program, gl.INFO_LOG_LENGTH, &logLength)

		// Fill linkage log with '\x00's. That's default in C strings
		linkLog := strings.Repeat("\x00", int(logLength+1))
		gl.GetProgramInfoLog(sp.program, logLength, nil, gl.Str(linkLog))

		return fmt.Errorf("failed to link program: %v", linkLog)
	}
	sp.linked = true

	return nil
}

func (sp *ShaderProgram) Use() {
	gl.UseProgram(sp.Program())
}

// Program returns OpenGL shader program ID
func (sp *ShaderProgram) Program() uint32 {
	if !sp.linked {
		return 0
	}
	return sp.program
}
