package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
)

type LeaveValidationResult struct {
	IsAllowed bool
	Reason    string
}

const (
	totalOfficeLeaveDays = 14
)

func canTakePersonalLeave(totalPublicLeaveDays int, joinDate, leaveStartDate time.Time, leaveDuration int) LeaveValidationResult {
	joinDatePlus180Days := joinDate.AddDate(0, 0, 180)
	if leaveStartDate.Before(joinDatePlus180Days) {
		return LeaveValidationResult{
			IsAllowed: false,
			Reason:    "Karena belum 180 hari sejak tanggal join karyawan.",
		}
	}

	personalLeaveQuota := totalOfficeLeaveDays - totalPublicLeaveDays

	yearEnd := time.Date(joinDatePlus180Days.Year(), 12, 31, 0, 0, 0, 0, time.UTC)
	daysRemaining := int(yearEnd.Sub(joinDatePlus180Days).Hours() / 24)

	personalLeaveQuotaForNewEmployee := (daysRemaining * personalLeaveQuota) / 365

	if leaveDuration > 3 {
		return LeaveValidationResult{
			IsAllowed: false,
			Reason:    "Cuti pribadi tidak boleh lebih dari 3 hari berturut-turut.",
		}
	}

	if leaveDuration > personalLeaveQuotaForNewEmployee {
		return LeaveValidationResult{
			IsAllowed: false,
			Reason:    fmt.Sprintf("Karena hanya boleh mengambil %d hari cuti.", personalLeaveQuotaForNewEmployee),
		}
	}

	return LeaveValidationResult{
		IsAllowed: true,
		Reason:    "Cuti pribadi diizinkan.",
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan jumlah cuti bersama:")
	scanner.Scan()
	totalPublicLeaveDays, _ := strconv.Atoi(scanner.Text())

	fmt.Println("Masukkan tanggal join karyawan (format: YYYY-MM-DD):")
	scanner.Scan()
	joinDate, _ := time.Parse("2006-01-02", scanner.Text())

	fmt.Println("Masukkan tanggal rencana cuti (format: YYYY-MM-DD):")
	scanner.Scan()
	leaveStartDate, _ := time.Parse("2006-01-02", scanner.Text())

	fmt.Println("Masukkan durasi cuti (hari):")
	scanner.Scan()
	leaveDuration, _ := strconv.Atoi(scanner.Text())

	result := canTakePersonalLeave(totalPublicLeaveDays, joinDate, leaveStartDate, leaveDuration)

	fmt.Printf("IsAllowed: %v\n", result.IsAllowed)
	fmt.Printf("Reason: %s\n", result.Reason)
}
