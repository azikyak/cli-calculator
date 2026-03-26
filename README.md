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
- macOS, Linux va Windows uchun build/install oqimi

## Talablar

- Go 1.20+
- `make` ixtiyoriy, lekin qulay
- Windows uchun PowerShell tavsiya qilinadi

## Loyiha Tuzilishi

```text
cli-calculator/
├── Makefile
├── README.md
├── go.mod
├── install.ps1
├── install.sh
├── main.go
└── .gitignore
```

## Ishga Tushirish

Loyiha papkasiga kiring:

```bash
cd cli-calculator
```

Windows PowerShell:

```powershell
cd .\cli-calculator
```

To'g'ridan-to'g'ri `go run` bilan:

```bash
go run . "2+3*4"
go run . "10 / (2 + 3)"
go run . "-(4+6)/2"
```

Interaktiv rejim:

```bash
go run .
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
macOS/Linux: bin/calc
Windows:     bin/calc.exe
```

Uni qo'lda ham ishlatish mumkin:

```bash
./bin/calc "7*(3-1)"
```

Windows PowerShell:

```powershell
.\bin\calc.exe "7*(3-1)"
```

## Install

Default install:

```bash
make install
```

OS bo'yicha default install joylari:

- macOS/Linux: `/usr/local/bin`, agar yozish huquqi bo'lmasa `~/.local/bin`
- Windows: `%USERPROFILE%\AppData\Local\Programs\cli-calculator\bin`

Unix shell orqali:

```bash
./install.sh
```

Windows PowerShell orqali:

```powershell
.\install.ps1
```

Agar boshqa joyga o'rnatmoqchi bo'lsangiz:

```bash
PREFIX=$HOME/.local/bin ./install.sh
```

yoki:

```bash
make install PREFIX=$HOME/.local/bin
```

Windows PowerShell:

```powershell
.\install.ps1 -Prefix "$HOME\AppData\Local\Microsoft\WindowsApps"
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
