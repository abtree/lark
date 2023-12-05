package main

import "lark/com/pkgs/xcfgstructpb"

func main() {
	xcfgstructpb.Run("../../configs", "../../pbfiles/auto_system.proto", true)
	xcfgstructpb.Run("../../apps/config/files", "../../pbfiles/auto_config.proto", false)
}
