package cmd

import (
	"github.com/spf13/cobra"
	//"log"
	"github.com/jukylin/esim/log"
	"github.com/jukylin/esim/tool/iface"
)

var ifaceCmd = &cobra.Command{
	Use:   "iface",
	Short: "根据接口生成空实例",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		writer := &iface.EsimWrite{}
		ifacer := iface.NewIface(writer)
		err := ifacer.Run(v)
		if err != nil {
			log.Log.Error(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(ifaceCmd)

	ifaceCmd.Flags().StringP("iname", "", "", "接口名称")

	ifaceCmd.Flags().StringP("out", "o", "", "输出文件: abc.go")

	ifaceCmd.Flags().StringP("ipath", "i", ".", "接口路径")

	ifaceCmd.Flags().BoolP("istar", "s", false, "带星")

	ifaceCmd.Flags().StringP("stname", "", "", "struct 名称：type struct_name struct{}")

	v.BindPFlags(ifaceCmd.Flags())
}
