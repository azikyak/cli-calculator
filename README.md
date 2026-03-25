# CLI Calculator

Oddiy, lekin amaliy terminal kalkulyator. Dastur matematik ifodani qabul qiladi, tokenlarga ajratadi, formatini tekshiradi va to'g'ri bo'lsa natijani hisoblaydi.

## Imkoniyatlar

- `+`, `-`, `*`, `/` operatorlari
- qavslar: `(` va `)`
- bo'shliq bilan yoki bo'shliqsiz ifodalar
- unary minus: `-5`, `-(2+3)`
- interaktiv rejim
- command line argument orqali ishlatish
- noto'g'ri input uchun tushunarli xatoliklar

## Talablar

- Go 1.20+
- `make` ixtiyoriy, lekin qulay

## Loyiha Tuzilishi

```text
cli-calculator/
├── Makefile
├── README.md
├── install.sh
├── main.go
└── .gitignore
```

## Ishga Tushirish

Loyiha papkasiga kiring:

```bash
cd cli-calculator
```

To'g'ridan-to'g'ri `go run` bilan:

```bash
go run main.go "2+3*4"
go run main.go "10 / (2 + 3)"
go run main.go "-(4+6)/2"
```

Interaktiv rejim:

```bash
go run main.go
```

Keyin terminalda ifoda kiriting:

```text
calc> 3*(2+5)
= 21
```

Chiqish uchun:

```text
exit
```

## Makefile Buyruqlari

```bash
make help
make build
make run ARGS='"2+3*4"'
make install
make uninstall
make clean
```

`make build` dan keyin binary shu yerda bo'ladi:

```text
bin/calc
```

Uni qo'lda ham ishlatish mumkin:

```bash
./bin/calc "7*(3-1)"
```

## Install

Default install:

```bash
make install
```

Agar `/usr/local/bin` ga yozish huquqi bo'lmasa, install script avtomatik ravishda `~/.local/bin` ga fallback qiladi.

Yoki script orqali:

```bash
./install.sh
```

Agar boshqa joyga o'rnatmoqchi bo'lsangiz:

```bash
PREFIX=$HOME/.local/bin ./install.sh
```

yoki:

```bash
make install PREFIX=$HOME/.local/bin
```

Agar `PATH` ichida bo'lsa, keyin shunday ishlaydi:

```bash
calc "12/(2+4)"
```

## Qo'llab-Quvvatlanadigan Ifodalar

To'g'ri misollar:

```text
2+3*4
10 / (2 + 3)
-8+5
-(4+6)/2
3.5*2
```

Xato misollar:

```text
2+*3
(2+3
4/0
abc
```
