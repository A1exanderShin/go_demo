package ifelse

func WeatherDef(isSunny, isWeekend bool) string {
	if isSunny && isWeekend {
		return ("Идеальный день для прогулки")
	} else if isSunny && !isWeekend {
		return ("Погода хорошая, но нужно пахать")
	} else if !isSunny && isWeekend {
		return ("Можно остаться дома и отдохнуть")
	} else {
		return ("Рабочий день с плохой погодой")
	}
}

func StatusStudent(mark int) string {
	switch mark {
	case 5:
		return ("Отличник что сказать")
	case 4:
		return ("Хорошняк пацан")
	case 3:
		return ("Будущий темщик")
	case 2:
		return ("Вор в законе")
	default:
		return ("Непонятки")
	}
}
