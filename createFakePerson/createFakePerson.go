package createFakePerson

import (
	uuid "github.com/google/uuid"
	faker "github.com/jaswdr/faker"
	"golang-fake-data/fakePersonStruct"
)

var f = faker.New()
var p = fakePersonStruct.Person{}

func CreatePerson() fakePersonStruct.Person {
	p.Id = uuid.New().String()
	p.Name = f.Person().Name()
	p.Gender = f.Person().Gender()
	p.Address = f.Address().Address()
	p.Country = f.Address().Country()
	p.City = f.Address().City()
	p.Continent = f.Address().CountryAbbr()
	p.Region = f.Address().CityPrefix()
	p.Age = int64(f.IntBetween(0, 100))
	p.Married = f.Bool()
	p.Phone = f.Phone().Number()
	p.CreditCardNumber = f.Payment().CreditCardNumber()
	p.CreditCardExpirationDateString = f.Payment().CreditCardExpirationDateString()
	p.CreditCardType = f.Payment().CreditCardType()
	p.TotalSpendingGold = f.Float64(1000, 1240, 999990)
	p.TotalSessionDuration = f.Float64(1000, 120, 999990)
	p.StartingDate = int64(f.IntBetween(1, 14))
	p.CurrentDate = int64(f.IntBetween(14, 31))
	p.StartingMonth = int64(f.IntBetween(10, 12))
	p.CurrentMonth = int64(f.IntBetween(10, 12))
	p.TotalClickEvent = int64(f.IntBetween(100, 100000))
	p.TotalSessionCount = int64(f.IntBetween(10, 1500))
	p.TotalScore = f.RandomFloat(10000, 1000, 10000)
	p.TotalGold = int64(f.IntBetween(1000, 95000))
	p.TotalSkillCount = int64(f.IntBetween(1000, 9500))
	p.SwipeLeftCount = int64(f.IntBetween(100, 500))
	p.SwipeRightCount = int64(f.IntBetween(100, 500))
	p.SwipeDownCount = int64(f.IntBetween(100, 500))
	p.SwipeUpCount = int64(f.IntBetween(100, 500))
	p.StartCor = f.RandomFloat(10000, 100, 1000)
	p.FinishCor = f.RandomFloat(10000, 100, 1000)
	p.RemainLife = f.RandomFloat(10000, -20, 100)
	p.AdvType = f.Car().Category()
	p.ProductType = f.Company().Name()
	p.DeviceType = int64(f.IntBetween(1, 80))
	p.GraphicsDeviceType = int64(f.IntBetween(134, 963))
	p.GraphicsMemorySize = int64(f.IntBetween(1000, 4000))
	p.OperatingSystem = int64(f.IntBetween(1, 5))
	p.ProcessorCount = int64(f.IntBetween(1, 32))
	p.ProcessorType = int64(f.IntBetween(1, 29))
	p.SystemMemorySize = int64(f.IntBetween(1000, 32000))
	p.FirstSessionYearOfDay = int64(f.IntBetween(1, 365))
	p.FirstSessionYear = int64(f.IntBetween(2020, 2022))
	p.FirstSessionWeekDay = int64(f.IntBetween(0, 6))
	p.FirstSessionHour = int64(f.IntBetween(0, 23))
	p.FirstSessionDuration = int64(f.IntBetween(0, 20))
	p.FirstSessionMinute = int64(f.IntBetween(1, 60))
	p.SecondSessionHour = int64(f.IntBetween(0, 2))
	p.SecondSessionDuration = int64(f.IntBetween(0, 12))
	p.SecondSessionMinute = int64(f.IntBetween(1, 60))
	p.ThirdSessionHour = int64(f.IntBetween(0, 2))
	p.ThirdSessionDuration = int64(f.IntBetween(0, 12))
	p.ThirdSessionMinute = int64(f.IntBetween(1, 60))
	p.FourthSessionHour = int64(f.IntBetween(0, 2))
	p.FourthSessionDuration = int64(f.IntBetween(0, 12))
	p.FourthSessionMinute = int64(f.IntBetween(1, 60))
	p.FifthSessionHour = int64(f.IntBetween(0, 2))
	p.FifthSessionDuration = int64(f.IntBetween(0, 12))
	p.FifthSessionMinute = int64(f.IntBetween(1, 60))
	p.PenultimateSessionHour = int64(f.IntBetween(0, 2))
	p.PenultimateSessionDuration = int64(f.IntBetween(0, 12))
	p.PenultimateSessionMinute = int64(f.IntBetween(1, 60))
	p.LastSessionYearOfDay = int64(f.IntBetween(1, 365))
	p.LastSessionYear = int64(f.IntBetween(2020, 2022))
	p.LastSessionHour = int64(f.IntBetween(0, 23))
	p.LastSessionDuration = int64(f.IntBetween(1, 30))
	p.LastSessionMinute = int64(f.IntBetween(1, 60))
	p.FirstHalfHourTotalSessionCount = int64(f.IntBetween(1, 5))
	p.FirstHalfHourTotalSessionDuration = int64(f.IntBetween(10, 60))
	p.FirstHourTotalSessionCount = int64(f.IntBetween(3, 9))
	p.FirstHourTotalSessionDuration = int64(f.IntBetween(20, 60))
	p.FirstTwoHourTotalSessionCount = int64(f.IntBetween(12, 90))
	p.FirstTwoHourTotalSessionDuration = int64(f.IntBetween(12, 98))
	p.FirstThreeHourTotalSessionCount = int64(f.IntBetween(16, 99))
	p.FirstThreeHourTotalSessionDuration = int64(f.IntBetween(16, 99))
	p.FirstSixHourTotalSessionCount = int64(f.IntBetween(16, 99))
	p.FirstSixHourTotalSessionDuration = int64(f.IntBetween(16, 99))
	p.FirstTwelveHourTotalSessionCount = int64(f.IntBetween(16, 99))
	p.FirstTwelveHourTotalSessionDuration = int64(f.IntBetween(16, 99))
	p.TotalSessionDay = int64(f.IntBetween(16, 99))
	p.TotalSessionHour = int64(f.IntBetween(16, 99))
	p.TotalSessionMinute = int64(f.IntBetween(16, 99))
	p.FirstDayTotalSessionCount = int64(f.IntBetween(3, 9))
	p.FirstDayTotalSessionDuration = int64(f.IntBetween(20, 60))
	p.SecondDayTotalSessionCount = int64(f.IntBetween(12, 90))
	p.SecondDayTotalSessionDuration = int64(f.IntBetween(12, 98))
	p.ThirdDayTotalSessionCount = int64(f.IntBetween(16, 99))
	p.ThirdDayTotalSessionDuration = int64(f.IntBetween(16, 99))
	p.FourthDayTotalSessionCount = int64(f.IntBetween(16, 99))
	p.FourthDayTotalSessionDuration = int64(f.IntBetween(16, 99))
	p.FifthDayTotalSessionCount = int64(f.IntBetween(16, 99))
	p.FifthDayTotalSessionDuration = int64(f.IntBetween(16, 99))
	p.SixthDayTotalSessionCount = int64(f.IntBetween(16, 99))
	p.SixthDayTotalSessionDuration = int64(f.IntBetween(16, 99))
	p.SeventhDayTotalSessionCount = int64(f.IntBetween(16, 99))
	p.SeventhDayTotalSessionDuration = int64(f.IntBetween(16, 99))
	p.MinSessionDuration = int64(f.IntBetween(16, 99))
	p.MaxSessionDuration = int64(f.IntBetween(16, 99))
	p.DailyAverageSessionCount = f.RandomFloat(10000, 4, 60)
	p.DailyAverageSessionDuration = f.RandomFloat(10000, 4, 60)
	p.SessionBasedAverageSessionDuration = f.RandomFloat(10000, 4, 60)
	p.DailyAverageSessionCountMinusFirstDaySessionCount = f.RandomFloat(10000, 4, 60)
	p.DailyAverageSessionDurationMinusFirstDaySessionDuration = f.RandomFloat(10000, 4, 60)
	p.SessionBasedAverageSessionDurationMinusFirstSessionDuration = f.RandomFloat(10000, 4, 60)
	p.SessionBasedAverageSessionDurationMinusLastSessionDuration = f.RandomFloat(10000, 4, 60)
	p.SundaySessionCount = int64(f.IntBetween(2, 22))
	p.MondaySessionCount = int64(f.IntBetween(2, 22))
	p.TuesdaySessionCount = int64(f.IntBetween(2, 22))
	p.WednesdaySessionCount = int64(f.IntBetween(2, 22))
	p.ThursdaySessionCount = int64(f.IntBetween(2, 22))
	p.FridaySessionCount = int64(f.IntBetween(2, 22))
	p.SaturdaySessionCount = int64(f.IntBetween(2, 22))
	p.AmSessionCount = int64(f.IntBetween(2, 22))
	p.PmSessionCount = int64(f.IntBetween(2, 22))
	p.SessionZeroToFiveHourCount = int64(f.IntBetween(2, 22))
	p.SessionSixToElevenHourCount = int64(f.IntBetween(2, 22))
	p.SessionTwelveToSeventeenHourCount = int64(f.IntBetween(2, 22))
	p.SessionEighteenToTwentyThreeHourCount = int64(f.IntBetween(2, 22))
	return p
}
