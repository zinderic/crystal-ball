package orb

type Moon struct {
	MoonPhase            string
	NorthernHemisphere   string
	SouthernHemisphere   string
	Visibility           string
	MidPhaseStandardTime string
	AverageMoonriseTime  string
	AverageMoonsetTime   string
}

var MoonData = []Moon{
	{
		MoonPhase:            "New Moon",
		NorthernHemisphere:   "disc completely in Sun's shadow (lit by earthshine only)",
		SouthernHemisphere:   "invisible (too close to Sun)",
		Visibility:           "noon",
		MidPhaseStandardTime: "6:00 AM",
		AverageMoonriseTime:  "6:00 PM",
		AverageMoonsetTime:   "none",
	},
	{
		MoonPhase:            "Waxing Crescent",
		NorthernHemisphere:   "right side, 0.1%–49.9% lit disc",
		SouthernHemisphere:   "left side, 0.1–49.9% lit disc",
		Visibility:           "late morning to post-dusk",
		MidPhaseStandardTime: "3:00 PM",
		AverageMoonriseTime:  "9:00 AM",
		AverageMoonsetTime:   "9:00 PM",
	},
	{
		MoonPhase:            "First Quarter",
		NorthernHemisphere:   "right side, 50% lit disc",
		SouthernHemisphere:   "left side, 50% lit disc",
		Visibility:           "afternoon and early evening",
		MidPhaseStandardTime: "6:00 PM",
		AverageMoonriseTime:  "noon",
		AverageMoonsetTime:   "midnight",
	},
	{
		MoonPhase:            "Waxing Gibbous",
		NorthernHemisphere:   "right side, 50.1%–99.9% lit disc",
		SouthernHemisphere:   "left side, 50.1%–99.9% lit disc",
		Visibility:           "late afternoon and most of night",
		MidPhaseStandardTime: "9:00 PM",
		AverageMoonriseTime:  "3:00 PM",
		AverageMoonsetTime:   "3:00 AM",
	},
	{
		MoonPhase:            "Full Moon",
		NorthernHemisphere:   "100% illuminated disc",
		SouthernHemisphere:   "sunset to sunrise (all night)",
		Visibility:           "midnight",
		MidPhaseStandardTime: "6:00 PM",
		AverageMoonriseTime:  "6:00 AM",
		AverageMoonsetTime:   "none",
	},
	{
		MoonPhase:            "Waning Gibbous",
		NorthernHemisphere:   "left side, 99.9%–50.1% lit disc",
		SouthernHemisphere:   "right side, 99.9%–50.1% lit disc",
		Visibility:           "most of night and early morning",
		MidPhaseStandardTime: "3:00 AM",
		AverageMoonriseTime:  "9:00 PM",
		AverageMoonsetTime:   "9:00 AM",
	},
	{
		MoonPhase:            "Last Quarter",
		NorthernHemisphere:   "left side, 50% lit disc",
		SouthernHemisphere:   "right side, 50% lit disc",
		Visibility:           "late night and morning",
		MidPhaseStandardTime: "6:00 AM",
		AverageMoonriseTime:  "midnight",
		AverageMoonsetTime:   "noon",
	},
	{
		MoonPhase:            "Waning Crescent",
		NorthernHemisphere:   "left side, 49.9%–0.1% lit disc",
		SouthernHemisphere:   "right side, 49.9%–0.1% lit disc",
		Visibility:           "pre-dawn to early afternoon",
		MidPhaseStandardTime: "9:00 AM",
		AverageMoonriseTime:  "3:00 AM",
		AverageMoonsetTime:   "3:00 PM",
	},
}
