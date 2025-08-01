package invest

import (
	"github.com/rykroon/fincli/internal/cli"
	"github.com/shopspring/decimal"
	"github.com/spf13/cobra"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

var fireNumberCmd = &cobra.Command{
	Use:   "fire-number",
	Short: "Calculate your FIRE number.",
	Run:   runFireNumberCmd,
}

type fireNumberFlags struct {
	AnnualExpenses    decimal.Decimal
	SafeWithdrawlRate decimal.Decimal
}

var fnf fireNumberFlags

func runFireNumberCmd(cmd *cobra.Command, args []string) {
	fmt := message.NewPrinter(language.English)
	fireNumber := fnf.AnnualExpenses.Div(fnf.SafeWithdrawlRate.Div(decimal.NewFromInt(100)))
	fmt.Printf("FIRE Number: $%v\n", fireNumber.StringFixed(2))
}

func init() {
	cli.AddDecimalVarP(fireNumberCmd.Flags(), &fnf.AnnualExpenses, "expenses", "e", "Annual expenses.")
	fnf.SafeWithdrawlRate = decimal.NewFromInt(4)
	cli.AddDecimalVar(fireNumberCmd.Flags(), &fnf.SafeWithdrawlRate, "swr", "Safe withdrawl rate.")

	fireNumberCmd.MarkFlagRequired("expenses")

	fireNumberCmd.Flags().SortFlags = false
	fireNumberCmd.Flags().PrintDefaults()
}
