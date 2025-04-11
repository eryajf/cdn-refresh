package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"cnb.cool/znb/doge-cdn-refresh/doge"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

func init() {
	// 首先从环境变量中获取配置
	viper.AutomaticEnv()
	viper.SetEnvPrefix("PLUGIN")
	viper.BindEnv("ak") // 获取环境变量 PLUGIN_AK 的值
	viper.BindEnv("sk")
	viper.BindEnv("rtype")
	viper.BindEnv("urls")
	rootCmd.Flags().StringP("ak", "a", viper.GetString("ak"), "Doge cloud access key [$PLUGIN_AK]")
	rootCmd.Flags().StringP("sk", "s", viper.GetString("sk"), "Doge cloud secret key [$PLUGIN_SK]")
	rootCmd.Flags().StringP("rtype", "t", viper.GetString("rtype"), "Refresh type (url/path) [$PLUGIN_RTYPE]")
	rootCmd.Flags().StringSliceP("urls", "u", strings.Split(viper.GetString("urls"), ","), "Refresh URLs [$PLUGIN_URLS]")
	// 遍历所有 flags，清除默认值占位符，避免在日志中打印
	rootCmd.Flags().VisitAll(func(f *pflag.Flag) {
		f.DefValue = ""
	})
}

var rootCmd = &cobra.Command{
	Use:     "doge-cdn-refresh",
	Short:   "Doge CDN refresh plugin",
	Version: "v0.1.0",
	RunE: func(cmd *cobra.Command, args []string) error {
		ak, _ := cmd.Flags().GetString("ak")
		sk, _ := cmd.Flags().GetString("sk")
		rtype, _ := cmd.Flags().GetString("rtype")
		urls, _ := cmd.Flags().GetStringSlice("urls")

		rst, err := doge.Refresh(ak, sk, rtype, urls)
		if err != nil {
			return err
		}
		if rst.Code != 200 {
			return fmt.Errorf("❌ refresh doge cdn failed")
		}
		fmt.Println("🎉 refresh doge cdn success")
		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
