package execute

import (
	"context"
	"main/schema"
	"os/exec"
)

func RunPython(ctx context.Context, code schema.ExecuteCode) ([]byte, error) {

	cmd := exec.CommandContext(ctx, "python3", "-c", *code.Code)

	return cmd.CombinedOutput()

	// output, err := cmd.CombinedOutput()
	// if err != nil {
	// 	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
	// 		utils.Log.Println("Code timed-out")
	// 	}
	// 	utils.Log.Println("Err exec-ing code", err, ctx.Err())
	// }

	// utils.Log.Printf("Code output: %s", output)
}
