package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"cnb.cool/znb/cdn-refresh/pkg/alidcdn"
	"cnb.cool/znb/cdn-refresh/pkg/aliesa"
	"cnb.cool/znb/cdn-refresh/pkg/doge"
	"cnb.cool/znb/cdn-refresh/pkg/qiniucdn"
	"cnb.cool/znb/cdn-refresh/pkg/tencentcdn"
	"cnb.cool/znb/cdn-refresh/pkg/tencenteo"
	"cnb.cool/znb/cdn-refresh/pkg/tools"
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
	viper.BindEnv("kind")
	viper.BindEnv("domain")
	viper.BindEnv("rtype")
	viper.BindEnv("urls")
	rootCmd.Flags().StringP("ak", "a", viper.GetString("ak"), "cloud access key [$PLUGIN_AK]")
	rootCmd.Flags().StringP("sk", "s", viper.GetString("sk"), "cloud secret key [$PLUGIN_SK]")
	rootCmd.Flags().StringP("kind", "k", viper.GetString("kind"), "cdn kind (doge/tencenteo/tencentcdn/aliesa/alidcdn/qiniucdn) [$PLUGIN_KIND]")
	rootCmd.Flags().StringP("domain", "d", viper.GetString("domain"), "domain name [$PLUGIN_DOMAIN]")
	rootCmd.Flags().StringP("rtype", "t", viper.GetString("rtype"), "Refresh type (url/path) [$PLUGIN_RTYPE]")
	rootCmd.Flags().StringSliceP("urls", "u", strings.Split(viper.GetString("urls"), ","), "Refresh URLs [$PLUGIN_URLS]")
	// 遍历所有 flags，清除默认值占位符，避免在日志中打印
	rootCmd.Flags().VisitAll(func(f *pflag.Flag) {
		f.DefValue = ""
	})
}

var rootCmd = &cobra.Command{
	Use:     "cdn-refresh",
	Short:   "🍒 Cloud CDN Refresh Plugin",
	Version: "v0.1.0",
	RunE: func(cmd *cobra.Command, args []string) error {
		ak, _ := cmd.Flags().GetString("ak")
		sk, _ := cmd.Flags().GetString("sk")
		kind, _ := cmd.Flags().GetString("kind")
		domain, _ := cmd.Flags().GetString("domain")
		rtype, _ := cmd.Flags().GetString("rtype")
		urls, _ := cmd.Flags().GetStringSlice("urls")

		// 如果没有提供任何参数或必要参数缺失，显示帮助信息
		if len(os.Args) == 1 && kind == "" {
			cmd.Help()
			return nil
		}

		r := tools.RefreshReq{
			Ak:       ak,
			Sk:       sk,
			ZoneName: domain,
			Rtype:    rtype,
			Urls:     urls,
		}
		switch kind {
		case "doge":
			rst, err := doge.Refresh(r)
			if err != nil || rst.Code != 200 {
				return fmt.Errorf("❌ refresh doge cdn failed")
			}
			fmt.Println("🎉 refresh doge cdn success")
		case "tencenteo":
			err := tencenteo.Refresh(r)
			if err != nil {
				return fmt.Errorf("❌ refresh tencent eo failed: %v", err)
			} else {
				fmt.Println("🎉 refresh tencent eo success")
			}
		case "tencentcdn":
			err := tencentcdn.Refresh(r)
			if err != nil {
				return fmt.Errorf("❌ refresh tencent cdn failed: %v", err)
			} else {
				fmt.Println("🎉 refresh tencent cdn success")
			}
		case "aliesa":
			err := aliesa.Refresh(r)
			if err != nil {
				return fmt.Errorf("❌ refresh ali esa failed: %v", err)
			} else {
				fmt.Println("🎉 refresh ali esa success")
			}
		case "alidcdn":
			err := alidcdn.Refresh(r)
			if err != nil {
				return fmt.Errorf("❌ refresh ali dcdn failed: %v", err)
			} else {
				fmt.Println("🎉 refresh ali dcdn success")
			}
		case "qiniucdn":
			err := qiniucdn.Refresh(r)
			if err != nil {
				return fmt.Errorf("❌ refresh qiniu cdn failed: %v", err)
			} else {
				fmt.Println("🎉 refresh qiniu cdn success")
			}
		}
		return nil
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
