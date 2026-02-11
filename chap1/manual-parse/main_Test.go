package main

// import (
// 	"bytes"
// 	"errors"
// 	"os"
// 	"os/exec"
// 	"strings"
// 	"testing"
// )


// func TestMain(m *testing.M){
// 	type tests_args []struct{
// 		args []string
// 		err  error
// 		exec	int
// 	}

// 	tests := []tests_args{
// 		args: []strings("-h"),
// 		err: nil,
// 		exec: os.Exit(1),

// 	}
// 	exitCode := m.Run()
//     os.Exit(exitCode)
// }



// 	err := os.Exit(0)
// 	for _, tc{
// 		if err := nil && os.Exit(1){
// 			err.Error
// 		}
// 	}
// 	exitCode := m.Run()
//     os.Exit(exitCode)
// }


// func TestMainSimple(t *testing.T) {
//     testCases := []struct {
//         name           string
//         args           []string
//         expectedCode   int
//         expectedStdout string
//         expectedStderr string
//     }{
//         {
//             name:           "No arguments",
//             args:           []string{},
//             expectedCode:   0,
//             expectedStdout: "No arguments provided\n",
//             expectedStderr: "",
//         },
//         {
//             name:           "Help flag",
//             args:           []string{"--help"},
//             expectedCode:   0,
//             expectedStdout: "Help message\n",
//             expectedStderr: "",
//         },
//         {
//             name:           "Error flag",
//             args:           []string{"--error"},
//             expectedCode:   1,
//             expectedStdout: "",
//             expectedStderr: "Error occurred",
//         },
//     }
    
//     for _, tc := range testCases {  
//             if exitCode != tc.expectedCode {
//                 t.Errorf("Exit code: expected %d, got %d", 
//                     tc.expectedCode, exitCode)
//             }
            
//             if stdout != tc.expectedStdout {
//                 t.Errorf("Stdout: expected %q, got %q", 
//                     tc.expectedStdout, stdout)
//             }
            
//             if stderr != tc.expectedStderr {
//                 t.Errorf("Stderr: expected %q, got %q", 
//                     tc.expectedStderr, stderr)
//             }
//         }
//     }
