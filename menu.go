
package cli

import (
	"bufio"
	"context"
	"fmt"
	"mini-soccer-cli/usecase"
	"os"
	"strings"
)

func Run(ctx context.Context, uc *usecase.ReportUsecase) {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("\n=== MINI SOCCER REPORT ===")
		fmt.Println("1. Revenue Report")
		fmt.Println("2. Pending Payments Report")
		fmt.Println("0. Exit")
		fmt.Print("Pilih menu: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			return
		}

		switch strings.TrimSpace(input) {
		case "1":
			items, err := uc.Revenue(ctx)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("\n=== REVENUE REPORT ===")
			for _, item := range items {
				fmt.Printf("Field: %s | Bookings: %d | Revenue: %.2f\n",
					item.FieldName,
					item.TotalBookings,
					item.TotalRevenue)
			}

		case "2":
			items, err := uc.PendingPayments(ctx)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}

			fmt.Println("\n=== PENDING PAYMENTS ===")
			if len(items) == 0 {
				fmt.Println("Tidak ada pembayaran tertunda.")
			}

			for _, item := range items {
				fmt.Printf("Customer: %s | Booking ID: %d | Date: %s\n",
					item.CustomerName,
					item.BookingID,
					item.BookingDate.Format("2006-01-02"))
			}

		case "0":
			fmt.Println("Selesai.")
			return

		default:
			fmt.Println("Menu tidak valid.")
		}
	}
}
