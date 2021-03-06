// Automatically generated file. DO NOT EDIT.
// Generated by gen-device.py from ATmega16U4.atdf, see http://packs.download.atmel.com/

// +build avr,atmega16u4

// Device information for the ATmega16U4.
//
package avr

// Magic type name for the compiler.
type __reg uint8

// Export this magic type name.
type RegValue = __reg

// Some information about this device.
const (
	DEVICE = "ATmega16U4"
	ARCH   = "AVR8"
	FAMILY = "megaAVR"
)

// Interrupts
const (
	IRQ_RESET        = 0  // External Pin,Power-on Reset,Brown-out Reset,Watchdog Reset,and JTAG AVR Reset. See Datasheet.
	IRQ_INT0         = 1  // External Interrupt Request 0
	IRQ_INT1         = 2  // External Interrupt Request 1
	IRQ_INT2         = 3  // External Interrupt Request 2
	IRQ_INT3         = 4  // External Interrupt Request 3
	IRQ_Reserved1    = 5  // Reserved1
	IRQ_Reserved2    = 6  // Reserved2
	IRQ_INT6         = 7  // External Interrupt Request 6
	IRQ_Reserved3    = 8  // Reserved3
	IRQ_PCINT0       = 9  // Pin Change Interrupt Request 0
	IRQ_USB_GEN      = 10 // USB General Interrupt Request
	IRQ_USB_COM      = 11 // USB Endpoint/Pipe Interrupt Communication Request
	IRQ_WDT          = 12 // Watchdog Time-out Interrupt
	IRQ_Reserved4    = 13 // Reserved4
	IRQ_Reserved5    = 14 // Reserved5
	IRQ_Reserved6    = 15 // Reserved6
	IRQ_TIMER1_CAPT  = 16 // Timer/Counter1 Capture Event
	IRQ_TIMER1_COMPA = 17 // Timer/Counter1 Compare Match A
	IRQ_TIMER1_COMPB = 18 // Timer/Counter1 Compare Match B
	IRQ_TIMER1_COMPC = 19 // Timer/Counter1 Compare Match C
	IRQ_TIMER1_OVF   = 20 // Timer/Counter1 Overflow
	IRQ_TIMER0_COMPA = 21 // Timer/Counter0 Compare Match A
	IRQ_TIMER0_COMPB = 22 // Timer/Counter0 Compare Match B
	IRQ_TIMER0_OVF   = 23 // Timer/Counter0 Overflow
	IRQ_SPI_STC      = 24 // SPI Serial Transfer Complete
	IRQ_USART1_RX    = 25 // USART1, Rx Complete
	IRQ_USART1_UDRE  = 26 // USART1 Data register Empty
	IRQ_USART1_TX    = 27 // USART1, Tx Complete
	IRQ_ANALOG_COMP  = 28 // Analog Comparator
	IRQ_ADC          = 29 // ADC Conversion Complete
	IRQ_EE_READY     = 30 // EEPROM Ready
	IRQ_TIMER3_CAPT  = 31 // Timer/Counter3 Capture Event
	IRQ_TIMER3_COMPA = 32 // Timer/Counter3 Compare Match A
	IRQ_TIMER3_COMPB = 33 // Timer/Counter3 Compare Match B
	IRQ_TIMER3_COMPC = 34 // Timer/Counter3 Compare Match C
	IRQ_TIMER3_OVF   = 35 // Timer/Counter3 Overflow
	IRQ_TWI          = 36 // 2-wire Serial Interface
	IRQ_SPM_READY    = 37 // Store Program Memory Read
	IRQ_TIMER4_COMPA = 38 // Timer/Counter4 Compare Match A
	IRQ_TIMER4_COMPB = 39 // Timer/Counter4 Compare Match B
	IRQ_TIMER4_COMPD = 40 // Timer/Counter4 Compare Match D
	IRQ_TIMER4_OVF   = 41 // Timer/Counter4 Overflow
	IRQ_TIMER4_FPF   = 42 // Timer/Counter4 Fault Protection Interrupt
	IRQ_max          = 42 // Highest interrupt number on this device.
)

// Peripherals
var (
	// Fuses
	FUSE = struct {
		EXTENDED __reg
		HIGH     __reg
		LOW      __reg
	}{
		EXTENDED: 0x2,
		HIGH:     0x1,
		LOW:      0x0,
	}

	// Lockbits
	LOCKBIT = struct {
		LOCKBIT __reg
	}{
		LOCKBIT: 0x0,
	}

	// Watchdog Timer
	WDT = struct {
		WDTCSR __reg
	}{
		WDTCSR: 0x60, // Watchdog Timer Control Register
	}

	// I/O Port
	PORT = struct {
		PORTD __reg
		DDRD  __reg
		PIND  __reg
		PORTB __reg
		DDRB  __reg
		PINB  __reg
		PORTC __reg
		DDRC  __reg
		PINC  __reg
		PORTE __reg
		DDRE  __reg
		PINE  __reg
		PORTF __reg
		DDRF  __reg
		PINF  __reg
	}{
		PORTD: 0x2b, // Port D Data Register
		DDRD:  0x2a, // Port D Data Direction Register
		PIND:  0x29, // Port D Input Pins
		PORTB: 0x25, // Port B Data Register
		DDRB:  0x24, // Port B Data Direction Register
		PINB:  0x23, // Port B Input Pins
		PORTC: 0x28, // Port C Data Register
		DDRC:  0x27, // Port C Data Direction Register
		PINC:  0x26, // Port C Input Pins
		PORTE: 0x2e, // Data Register, Port E
		DDRE:  0x2d, // Data Direction Register, Port E
		PINE:  0x2c, // Input Pins, Port E
		PORTF: 0x31, // Data Register, Port F
		DDRF:  0x30, // Data Direction Register, Port F
		PINF:  0x2f, // Input Pins, Port F
	}

	// Serial Peripheral Interface
	SPI = struct {
		SPCR __reg
		SPSR __reg
		SPDR __reg
	}{
		SPCR: 0x4c, // SPI Control Register
		SPSR: 0x4d, // SPI Status Register
		SPDR: 0x4e, // SPI Data Register
	}

	// USART
	USART = struct {
		UDR1   __reg
		UCSR1A __reg
		UCSR1B __reg
		UCSR1C __reg
		UCSR1D __reg
		UBRR1L __reg
		UBRR1H __reg
	}{
		UDR1:   0xce, // USART I/O Data Register
		UCSR1A: 0xc8, // USART Control and Status Register A
		UCSR1B: 0xc9, // USART Control and Status Register B
		UCSR1C: 0xca, // USART Control and Status Register C
		UCSR1D: 0xcb, // USART Control and Status Register D
		UBRR1L: 0xcc, // USART Baud Rate Register  Bytes
		UBRR1H: 0xcc, // USART Baud Rate Register  Bytes
	}

	// Bootloader
	BOOT_LOAD = struct {
		SPMCSR __reg
	}{
		SPMCSR: 0x57, // Store Program Memory Control Register
	}

	// EEPROM
	EEPROM = struct {
		EEARL __reg
		EEARH __reg
		EEDR  __reg
		EECR  __reg
	}{
		EEARL: 0x41, // EEPROM Address Register Low Bytes
		EEARH: 0x41, // EEPROM Address Register Low Bytes
		EEDR:  0x40, // EEPROM Data Register
		EECR:  0x3f, // EEPROM Control Register
	}

	// Timer/Counter, 8-bit
	TC8 = struct {
		OCR0B  __reg
		OCR0A  __reg
		TCNT0  __reg
		TCCR0B __reg
		TCCR0A __reg
		TIMSK0 __reg
		TIFR0  __reg
		GTCCR  __reg
	}{
		OCR0B:  0x48, // Timer/Counter0 Output Compare Register
		OCR0A:  0x47, // Timer/Counter0 Output Compare Register
		TCNT0:  0x46, // Timer/Counter0
		TCCR0B: 0x45, // Timer/Counter Control Register B
		TCCR0A: 0x44, // Timer/Counter  Control Register A
		TIMSK0: 0x6e, // Timer/Counter0 Interrupt Mask Register
		TIFR0:  0x35, // Timer/Counter0 Interrupt Flag register
		GTCCR:  0x43, // General Timer/Counter Control Register
	}

	// Timer/Counter, 10-bit
	TC10 = struct {
		TCCR4A __reg
		TCCR4B __reg
		TCCR4C __reg
		TCCR4D __reg
		TCCR4E __reg
		TCNT4  __reg
		TC4H   __reg
		OCR4A  __reg
		OCR4B  __reg
		OCR4C  __reg
		OCR4D  __reg
		TIMSK4 __reg
		TIFR4  __reg
		DT4    __reg
	}{
		TCCR4A: 0xc0, // Timer/Counter4 Control Register A
		TCCR4B: 0xc1, // Timer/Counter4 Control Register B
		TCCR4C: 0xc2, // Timer/Counter 4 Control Register C
		TCCR4D: 0xc3, // Timer/Counter 4 Control Register D
		TCCR4E: 0xc4, // Timer/Counter 4 Control Register E
		TCNT4:  0xbe, // Timer/Counter4 Low Bytes
		TC4H:   0xbf, // Timer/Counter4
		OCR4A:  0xcf, // Timer/Counter4 Output Compare Register A
		OCR4B:  0xd0, // Timer/Counter4 Output Compare Register B
		OCR4C:  0xd1, // Timer/Counter4 Output Compare Register C
		OCR4D:  0xd2, // Timer/Counter4 Output Compare Register D
		TIMSK4: 0x72, // Timer/Counter4 Interrupt Mask Register
		TIFR4:  0x39, // Timer/Counter4 Interrupt Flag register
		DT4:    0xd4, // Timer/Counter 4 Dead Time Value
	}

	// Timer/Counter, 16-bit
	TC16 = struct {
		TCCR3A __reg
		TCCR3B __reg
		TCCR3C __reg
		TCNT3L __reg
		TCNT3H __reg
		OCR3AL __reg
		OCR3AH __reg
		OCR3BL __reg
		OCR3BH __reg
		OCR3CL __reg
		OCR3CH __reg
		ICR3L  __reg
		ICR3H  __reg
		TIMSK3 __reg
		TIFR3  __reg
		TCCR1A __reg
		TCCR1B __reg
		TCCR1C __reg
		TCNT1L __reg
		TCNT1H __reg
		OCR1AL __reg
		OCR1AH __reg
		OCR1BL __reg
		OCR1BH __reg
		OCR1CL __reg
		OCR1CH __reg
		ICR1L  __reg
		ICR1H  __reg
		TIMSK1 __reg
		TIFR1  __reg
	}{
		TCCR3A: 0x90, // Timer/Counter3 Control Register A
		TCCR3B: 0x91, // Timer/Counter3 Control Register B
		TCCR3C: 0x92, // Timer/Counter 3 Control Register C
		TCNT3L: 0x94, // Timer/Counter3 Bytes
		TCNT3H: 0x94, // Timer/Counter3 Bytes
		OCR3AL: 0x98, // Timer/Counter3 Output Compare Register A  Bytes
		OCR3AH: 0x98, // Timer/Counter3 Output Compare Register A  Bytes
		OCR3BL: 0x9a, // Timer/Counter3 Output Compare Register B Bytes
		OCR3BH: 0x9a, // Timer/Counter3 Output Compare Register B Bytes
		OCR3CL: 0x9c, // Timer/Counter3 Output Compare Register C Bytes
		OCR3CH: 0x9c, // Timer/Counter3 Output Compare Register C Bytes
		ICR3L:  0x96, // Timer/Counter3 Input Capture Register  Bytes
		ICR3H:  0x96, // Timer/Counter3 Input Capture Register  Bytes
		TIMSK3: 0x71, // Timer/Counter3 Interrupt Mask Register
		TIFR3:  0x38, // Timer/Counter3 Interrupt Flag register
		TCCR1A: 0x80, // Timer/Counter1 Control Register A
		TCCR1B: 0x81, // Timer/Counter1 Control Register B
		TCCR1C: 0x82, // Timer/Counter 1 Control Register C
		TCNT1L: 0x84, // Timer/Counter1 Bytes
		TCNT1H: 0x84, // Timer/Counter1 Bytes
		OCR1AL: 0x88, // Timer/Counter1 Output Compare Register A Bytes
		OCR1AH: 0x88, // Timer/Counter1 Output Compare Register A Bytes
		OCR1BL: 0x8a, // Timer/Counter1 Output Compare Register B Bytes
		OCR1BH: 0x8a, // Timer/Counter1 Output Compare Register B Bytes
		OCR1CL: 0x8c, // Timer/Counter1 Output Compare Register C Bytes
		OCR1CH: 0x8c, // Timer/Counter1 Output Compare Register C Bytes
		ICR1L:  0x86, // Timer/Counter1 Input Capture Register Bytes
		ICR1H:  0x86, // Timer/Counter1 Input Capture Register Bytes
		TIMSK1: 0x6f, // Timer/Counter1 Interrupt Mask Register
		TIFR1:  0x36, // Timer/Counter1 Interrupt Flag register
	}

	// JTAG Interface
	JTAG = struct {
		OCDR __reg
	}{
		OCDR: 0x51, // On-Chip Debug Related Register in I/O Memory
	}

	// External Interrupts
	EXINT = struct {
		EICRA  __reg
		EICRB  __reg
		EIMSK  __reg
		EIFR   __reg
		PCMSK0 __reg
		PCIFR  __reg
		PCICR  __reg
	}{
		EICRA:  0x69, // External Interrupt Control Register A
		EICRB:  0x6a, // External Interrupt Control Register B
		EIMSK:  0x3d, // External Interrupt Mask Register
		EIFR:   0x3c, // External Interrupt Flag Register
		PCMSK0: 0x6b, // Pin Change Mask Register 0
		PCIFR:  0x3b, // Pin Change Interrupt Flag Register
		PCICR:  0x68, // Pin Change Interrupt Control Register
	}

	// Two Wire Serial Interface
	TWI = struct {
		TWAMR __reg
		TWBR  __reg
		TWCR  __reg
		TWSR  __reg
		TWDR  __reg
		TWAR  __reg
	}{
		TWAMR: 0xbd, // TWI (Slave) Address Mask Register
		TWBR:  0xb8, // TWI Bit Rate register
		TWCR:  0xbc, // TWI Control Register
		TWSR:  0xb9, // TWI Status Register
		TWDR:  0xbb, // TWI Data register
		TWAR:  0xba, // TWI (Slave) Address register
	}

	// Analog-to-Digital Converter
	ADC = struct {
		ADMUX  __reg
		ADCSRA __reg
		ADCL   __reg
		ADCH   __reg
		DIDR0  __reg
		DIDR2  __reg
	}{
		ADMUX:  0x7c, // The ADC multiplexer Selection Register
		ADCSRA: 0x7a, // The ADC Control and Status register
		ADCL:   0x78, // ADC Data Register  Bytes
		ADCH:   0x78, // ADC Data Register  Bytes
		DIDR0:  0x7e, // Digital Input Disable Register 1
		DIDR2:  0x7d, // Digital Input Disable Register 1
	}

	// Analog Comparator
	AC = struct {
		ACSR  __reg
		DIDR1 __reg
	}{
		ACSR:  0x50, // Analog Comparator Control And Status Register
		DIDR1: 0x7f,
	}

	// CPU Registers
	CPU = struct {
		SREG    __reg
		SPL     __reg
		SPH     __reg
		OSCCAL  __reg
		RCCTRL  __reg
		CLKPR   __reg
		SMCR    __reg
		EIND    __reg
		RAMPZ   __reg
		GPIOR2  __reg
		GPIOR1  __reg
		GPIOR0  __reg
		PRR1    __reg
		PRR0    __reg
		CLKSTA  __reg
		CLKSEL1 __reg
		CLKSEL0 __reg
	}{
		SREG:    0x5f, // Status Register
		SPL:     0x5d, // Stack Pointer
		SPH:     0x5d, // Stack Pointer
		OSCCAL:  0x66, // Oscillator Calibration Value
		RCCTRL:  0x67, // Oscillator Control Register
		CLKPR:   0x61,
		SMCR:    0x53, // Sleep Mode Control Register
		EIND:    0x5c, // Extended Indirect Register
		RAMPZ:   0x5b, // Extended Z-pointer Register for ELPM/SPM
		GPIOR2:  0x4b, // General Purpose IO Register 2
		GPIOR1:  0x4a, // General Purpose IO Register 1
		GPIOR0:  0x3e, // General Purpose IO Register 0
		PRR1:    0x65, // Power Reduction Register1
		PRR0:    0x64, // Power Reduction Register0
		CLKSTA:  0xc7,
		CLKSEL1: 0xc6,
		CLKSEL0: 0xc5,
	}

	// Phase Locked Loop
	PLL = struct {
		PLLCSR __reg
		PLLFRQ __reg
	}{
		PLLCSR: 0x49, // PLL Status and Control register
		PLLFRQ: 0x52, // PLL Frequency Control Register
	}

	// USB Device Registers
	USB_DEVICE = struct {
		UEINT   __reg
		UEBCHX  __reg
		UEBCLX  __reg
		UEDATX  __reg
		UEIENX  __reg
		UESTA1X __reg
		UESTA0X __reg
		UECFG1X __reg
		UECFG0X __reg
		UECONX  __reg
		UERST   __reg
		UENUM   __reg
		UEINTX  __reg
		UDMFN   __reg
		UDFNUML __reg
		UDFNUMH __reg
		UDADDR  __reg
		UDIEN   __reg
		UDINT   __reg
		UDCON   __reg
		USBCON  __reg
		USBINT  __reg
		USBSTA  __reg
		UHWCON  __reg
	}{
		UEINT:   0xf4,
		UEBCHX:  0xf3,
		UEBCLX:  0xf2,
		UEDATX:  0xf1,
		UEIENX:  0xf0,
		UESTA1X: 0xef,
		UESTA0X: 0xee,
		UECFG1X: 0xed,
		UECFG0X: 0xec,
		UECONX:  0xeb,
		UERST:   0xea,
		UENUM:   0xe9,
		UEINTX:  0xe8,
		UDMFN:   0xe6,
		UDFNUML: 0xe4,
		UDFNUMH: 0xe4,
		UDADDR:  0xe3,
		UDIEN:   0xe2,
		UDINT:   0xe1,
		UDCON:   0xe0,
		USBCON:  0xd8, // USB General Control Register
		USBINT:  0xda,
		USBSTA:  0xd9,
		UHWCON:  0xd7,
	}
)

// Bitfields for FUSE: Fuses
const (
	// EXTENDED
	EXTENDED_BODLEVEL = 0x7 // Brown-out Detector trigger level
	EXTENDED_HWBE     = 0x8 // Hardware Boot Enable

	// HIGH
	HIGH_OCDEN   = 0x80 // On-Chip Debug Enabled
	HIGH_JTAGEN  = 0x40 // JTAG Interface Enabled
	HIGH_SPIEN   = 0x20 // Serial program downloading (SPI) enabled
	HIGH_WDTON   = 0x10 // Watchdog timer always on
	HIGH_EESAVE  = 0x8  // Preserve EEPROM through the Chip Erase cycle
	HIGH_BOOTSZ  = 0x6  // Select Boot Size
	HIGH_BOOTRST = 0x1  // Boot Reset vector Enabled

	// LOW
	LOW_CKDIV8    = 0x80 // Divide clock by 8 internally
	LOW_CKOUT     = 0x40 // Clock output on PORTC7
	LOW_SUT_CKSEL = 0x3f // Select Clock Source
)

// Bitfields for LOCKBIT: Lockbits
const (
	// LOCKBIT
	LOCKBIT_LB   = 0x3  // Memory Lock
	LOCKBIT_BLB0 = 0xc  // Boot Loader Protection Mode
	LOCKBIT_BLB1 = 0x30 // Boot Loader Protection Mode
)

// Bitfields for WDT: Watchdog Timer
const (
	// WDTCSR: Watchdog Timer Control Register
	WDTCSR_WDIF = 0x80 // Watchdog Timeout Interrupt Flag
	WDTCSR_WDIE = 0x40 // Watchdog Timeout Interrupt Enable
	WDTCSR_WDP  = 0x27 // Watchdog Timer Prescaler Bits
	WDTCSR_WDCE = 0x10 // Watchdog Change Enable
	WDTCSR_WDE  = 0x8  // Watch Dog Enable
)

// Bitfields for SPI: Serial Peripheral Interface
const (
	// SPCR: SPI Control Register
	SPCR_SPIE = 0x80 // SPI Interrupt Enable
	SPCR_SPE  = 0x40 // SPI Enable
	SPCR_DORD = 0x20 // Data Order
	SPCR_MSTR = 0x10 // Master/Slave Select
	SPCR_CPOL = 0x8  // Clock polarity
	SPCR_CPHA = 0x4  // Clock Phase
	SPCR_SPR  = 0x3  // SPI Clock Rate Selects

	// SPSR: SPI Status Register
	SPSR_SPIF  = 0x80 // SPI Interrupt Flag
	SPSR_WCOL  = 0x40 // Write Collision Flag
	SPSR_SPI2X = 0x1  // Double SPI Speed Bit

	// SPDR: SPI Data Register
	SPDR_SPDR = 0xff // SPI Data bits
)

// Bitfields for USART: USART
const (
	// UDR1: USART I/O Data Register
	UDR1_UDR1 = 0xff // USART I/O Data bits

	// UCSR1A: USART Control and Status Register A
	UCSR1A_RXC1  = 0x80 // USART Receive Complete
	UCSR1A_TXC1  = 0x40 // USART Transmitt Complete
	UCSR1A_UDRE1 = 0x20 // USART Data Register Empty
	UCSR1A_FE1   = 0x10 // Framing Error
	UCSR1A_DOR1  = 0x8  // Data overRun
	UCSR1A_UPE1  = 0x4  // Parity Error
	UCSR1A_U2X1  = 0x2  // Double the USART transmission speed
	UCSR1A_MPCM1 = 0x1  // Multi-processor Communication Mode

	// UCSR1B: USART Control and Status Register B
	UCSR1B_RXCIE1 = 0x80 // RX Complete Interrupt Enable
	UCSR1B_TXCIE1 = 0x40 // TX Complete Interrupt Enable
	UCSR1B_UDRIE1 = 0x20 // USART Data register Empty Interrupt Enable
	UCSR1B_RXEN1  = 0x10 // Receiver Enable
	UCSR1B_TXEN1  = 0x8  // Transmitter Enable
	UCSR1B_UCSZ12 = 0x4  // Character Size
	UCSR1B_RXB81  = 0x2  // Receive Data Bit 8
	UCSR1B_TXB81  = 0x1  // Transmit Data Bit 8

	// UCSR1C: USART Control and Status Register C
	UCSR1C_UMSEL1 = 0xc0 // USART Mode Select
	UCSR1C_UPM1   = 0x30 // Parity Mode Bits
	UCSR1C_USBS1  = 0x8  // Stop Bit Select
	UCSR1C_UCSZ1  = 0x6  // Character Size
	UCSR1C_UCPOL1 = 0x1  // Clock Polarity

	// UCSR1D: USART Control and Status Register D
	UCSR1D_CTSEN = 0x2 // CTS Enable
	UCSR1D_RTSEN = 0x1 // RTS Enable
)

// Bitfields for BOOT_LOAD: Bootloader
const (
	// SPMCSR: Store Program Memory Control Register
	SPMCSR_SPMIE  = 0x80 // SPM Interrupt Enable
	SPMCSR_RWWSB  = 0x40 // Read While Write Section Busy
	SPMCSR_SIGRD  = 0x20 // Signature Row Read
	SPMCSR_RWWSRE = 0x10 // Read While Write section read enable
	SPMCSR_BLBSET = 0x8  // Boot Lock Bit Set
	SPMCSR_PGWRT  = 0x4  // Page Write
	SPMCSR_PGERS  = 0x2  // Page Erase
	SPMCSR_SPMEN  = 0x1  // Store Program Memory Enable
)

// Bitfields for EEPROM: EEPROM
const (
	// EEARL: EEPROM Address Register Low Bytes

	// EEARH: EEPROM Address Register Low Bytes
	EEAR_EEAR = 0xfff // EEPROM Address Bits

	// EEDR: EEPROM Data Register
	EEDR_EEDR = 0xff // EEPROM Data Bits

	// EECR: EEPROM Control Register
	EECR_EEPM  = 0x30 // EEPROM Programming Mode Bits
	EECR_EERIE = 0x8  // EEPROM Ready Interrupt Enable
	EECR_EEMPE = 0x4  // EEPROM Master Write Enable
	EECR_EEPE  = 0x2  // EEPROM Write Enable
	EECR_EERE  = 0x1  // EEPROM Read Enable
)

// Bitfields for TC8: Timer/Counter, 8-bit
const (
	// OCR0B: Timer/Counter0 Output Compare Register
	OCR0B_OCR0B = 0xff // Timer/Counter0 Output Compare B bits

	// OCR0A: Timer/Counter0 Output Compare Register
	OCR0A_OCR0A = 0xff // Timer/Counter0 Output Compare A bits

	// TCNT0: Timer/Counter0
	TCNT0_TCNT0 = 0xff // Timer/Counter0 bits

	// TCCR0B: Timer/Counter Control Register B
	TCCR0B_FOC0A = 0x80 // Force Output Compare A
	TCCR0B_FOC0B = 0x40 // Force Output Compare B
	TCCR0B_WGM02 = 0x8
	TCCR0B_CS0   = 0x7 // Clock Select

	// TCCR0A: Timer/Counter  Control Register A
	TCCR0A_COM0A = 0xc0 // Compare Output Mode, Phase Correct PWM Mode
	TCCR0A_COM0B = 0x30 // Compare Output Mode, Fast PWm
	TCCR0A_WGM0  = 0x3  // Waveform Generation Mode

	// TIMSK0: Timer/Counter0 Interrupt Mask Register
	TIMSK0_OCIE0B = 0x4 // Timer/Counter0 Output Compare Match B Interrupt Enable
	TIMSK0_OCIE0A = 0x2 // Timer/Counter0 Output Compare Match A Interrupt Enable
	TIMSK0_TOIE0  = 0x1 // Timer/Counter0 Overflow Interrupt Enable

	// TIFR0: Timer/Counter0 Interrupt Flag register
	TIFR0_OCF0B = 0x4 // Timer/Counter0 Output Compare Flag 0B
	TIFR0_OCF0A = 0x2 // Timer/Counter0 Output Compare Flag 0A
	TIFR0_TOV0  = 0x1 // Timer/Counter0 Overflow Flag

	// GTCCR: General Timer/Counter Control Register
	GTCCR_TSM     = 0x80 // Timer/Counter Synchronization Mode
	GTCCR_PSRSYNC = 0x1  // Prescaler Reset Timer/Counter1 and Timer/Counter0
)

// Bitfields for TC10: Timer/Counter, 10-bit
const (
	// TCCR4A: Timer/Counter4 Control Register A
	TCCR4A_COM4A = 0xc0 // Compare Output Mode 1A, bits
	TCCR4A_COM4B = 0x30 // Compare Output Mode 4B, bits
	TCCR4A_FOC4A = 0x8  // Force Output Compare Match 4A
	TCCR4A_FOC4B = 0x4  // Force Output Compare Match 4B
	TCCR4A_PWM4A = 0x2
	TCCR4A_PWM4B = 0x1

	// TCCR4B: Timer/Counter4 Control Register B
	TCCR4B_PWM4X = 0x80 // PWM Inversion Mode
	TCCR4B_PSR4  = 0x40 // Prescaler Reset Timer/Counter 4
	TCCR4B_DTPS4 = 0x30 // Dead Time Prescaler Bits
	TCCR4B_CS4   = 0xf  // Clock Select Bits

	// TCCR4C: Timer/Counter 4 Control Register C
	TCCR4C_COM4A1S = 0x80 // Comparator A Output Mode
	TCCR4C_COM4A0S = 0x40 // Comparator A Output Mode
	TCCR4C_COM4B1S = 0x20 // Comparator B Output Mode
	TCCR4C_COM4B0S = 0x10 // Comparator B Output Mode
	TCCR4C_COM4D   = 0xc  // Comparator D Output Mode
	TCCR4C_FOC4D   = 0x2  // Force Output Compare Match 4D
	TCCR4C_PWM4D   = 0x1  // Pulse Width Modulator D Enable

	// TCCR4D: Timer/Counter 4 Control Register D
	TCCR4D_FPIE4 = 0x80 // Fault Protection Interrupt Enable
	TCCR4D_FPEN4 = 0x40 // Fault Protection Mode Enable
	TCCR4D_FPNC4 = 0x20 // Fault Protection Noise Canceler
	TCCR4D_FPES4 = 0x10 // Fault Protection Edge Select
	TCCR4D_FPAC4 = 0x8  // Fault Protection Analog Comparator Enable
	TCCR4D_FPF4  = 0x4  // Fault Protection Interrupt Flag
	TCCR4D_WGM4  = 0x3  // Waveform Generation Mode bits

	// TCCR4E: Timer/Counter 4 Control Register E
	TCCR4E_TLOCK4 = 0x80 // Register Update Lock
	TCCR4E_ENHC4  = 0x40 // Enhanced Compare/PWM Mode
	TCCR4E_OC4OE  = 0x3f // Output Compare Override Enable bit

	// TCNT4: Timer/Counter4 Low Bytes
	TCNT4_TC4 = 0xff // Timer/Counter4 bits

	// TC4H: Timer/Counter4
	TC4H_TC4 = 0x7 // Timer/Counter4 bits

	// OCR4A: Timer/Counter4 Output Compare Register A
	OCR4A_OCR4A = 0xff // Timer/Counter4 Output Compare A bits

	// OCR4B: Timer/Counter4 Output Compare Register B
	OCR4B_OCR4B = 0xff // Timer/Counter4 Output Compare B bits

	// OCR4C: Timer/Counter4 Output Compare Register C
	OCR4C_OCR4C = 0xff // Timer/Counter4 Output Compare C bits

	// OCR4D: Timer/Counter4 Output Compare Register D
	OCR4D_OCR4D = 0xff // Timer/Counter4 Output Compare D bits

	// TIMSK4: Timer/Counter4 Interrupt Mask Register
	TIMSK4_OCIE4D = 0x80 // Timer/Counter4 Output Compare D Match Interrupt Enable
	TIMSK4_OCIE4A = 0x40 // Timer/Counter4 Output Compare A Match Interrupt Enable
	TIMSK4_OCIE4B = 0x20 // Timer/Counter4 Output Compare B Match Interrupt Enable
	TIMSK4_TOIE4  = 0x4  // Timer/Counter4 Overflow Interrupt Enable

	// TIFR4: Timer/Counter4 Interrupt Flag register
	TIFR4_OCF4D = 0x80 // Output Compare Flag 4D
	TIFR4_OCF4A = 0x40 // Output Compare Flag 4A
	TIFR4_OCF4B = 0x20 // Output Compare Flag 4B
	TIFR4_TOV4  = 0x4  // Timer/Counter4 Overflow Flag

	// DT4: Timer/Counter 4 Dead Time Value
	DT4_DT4L = 0xff // Timer/Counter 4 Dead Time Value Bits
)

// Bitfields for TC16: Timer/Counter, 16-bit
const (
	// TCCR3A: Timer/Counter3 Control Register A
	TCCR3A_COM3A = 0xc0 // Compare Output Mode 1A, bits
	TCCR3A_COM3B = 0x30 // Compare Output Mode 3B, bits
	TCCR3A_COM3C = 0xc  // Compare Output Mode 3C, bits
	TCCR3A_WGM3  = 0x3  // Waveform Generation Mode

	// TCCR3B: Timer/Counter3 Control Register B
	TCCR3B_ICNC3 = 0x80 // Input Capture 3 Noise Canceler
	TCCR3B_ICES3 = 0x40 // Input Capture 3 Edge Select
	TCCR3B_WGM3  = 0x18 // Waveform Generation Mode
	TCCR3B_CS3   = 0x7  // Prescaler source of Timer/Counter 3

	// TCCR3C: Timer/Counter 3 Control Register C
	TCCR3C_FOC3A = 0x80 // Force Output Compare 3A
	TCCR3C_FOC3B = 0x40 // Force Output Compare 3B
	TCCR3C_FOC3C = 0x20 // Force Output Compare 3C

	// TCNT3L: Timer/Counter3 Bytes

	// TCNT3H: Timer/Counter3 Bytes
	TCNT3_TCNT3 = 0xffff // Timer/Counter3 bits

	// OCR3AL: Timer/Counter3 Output Compare Register A  Bytes

	// OCR3AH: Timer/Counter3 Output Compare Register A  Bytes
	OCR3A_OCR3A = 0xffff // Timer/Counter3 Output Compare A bits

	// OCR3BL: Timer/Counter3 Output Compare Register B Bytes

	// OCR3BH: Timer/Counter3 Output Compare Register B Bytes
	OCR3B_OCR3B = 0xffff // Timer/Counter3 Output Compare B bits

	// OCR3CL: Timer/Counter3 Output Compare Register C Bytes

	// OCR3CH: Timer/Counter3 Output Compare Register C Bytes
	OCR3C_OCR3C = 0xffff // Timer/Counter3 Output Compare C bits

	// ICR3L: Timer/Counter3 Input Capture Register  Bytes

	// ICR3H: Timer/Counter3 Input Capture Register  Bytes
	ICR3_ICR3 = 0xffff // Timer/Counter3 Input Capture bits

	// TIMSK3: Timer/Counter3 Interrupt Mask Register
	TIMSK3_ICIE3  = 0x20 // Timer/Counter3 Input Capture Interrupt Enable
	TIMSK3_OCIE3C = 0x8  // Timer/Counter3 Output Compare C Match Interrupt Enable
	TIMSK3_OCIE3B = 0x4  // Timer/Counter3 Output Compare B Match Interrupt Enable
	TIMSK3_OCIE3A = 0x2  // Timer/Counter3 Output Compare A Match Interrupt Enable
	TIMSK3_TOIE3  = 0x1  // Timer/Counter3 Overflow Interrupt Enable

	// TIFR3: Timer/Counter3 Interrupt Flag register
	TIFR3_ICF3  = 0x20 // Input Capture Flag 3
	TIFR3_OCF3C = 0x8  // Output Compare Flag 3C
	TIFR3_OCF3B = 0x4  // Output Compare Flag 3B
	TIFR3_OCF3A = 0x2  // Output Compare Flag 3A
	TIFR3_TOV3  = 0x1  // Timer/Counter3 Overflow Flag

	// TCCR1A: Timer/Counter1 Control Register A
	TCCR1A_COM1A = 0xc0 // Compare Output Mode 1A, bits
	TCCR1A_COM1B = 0x30 // Compare Output Mode 1B, bits
	TCCR1A_COM1C = 0xc  // Compare Output Mode 1C, bits
	TCCR1A_WGM1  = 0x3  // Waveform Generation Mode

	// TCCR1B: Timer/Counter1 Control Register B
	TCCR1B_ICNC1 = 0x80 // Input Capture 1 Noise Canceler
	TCCR1B_ICES1 = 0x40 // Input Capture 1 Edge Select
	TCCR1B_WGM1  = 0x18 // Waveform Generation Mode
	TCCR1B_CS1   = 0x7  // Prescaler source of Timer/Counter 1

	// TCCR1C: Timer/Counter 1 Control Register C
	TCCR1C_FOC1A = 0x80 // Force Output Compare 1A
	TCCR1C_FOC1B = 0x40 // Force Output Compare 1B
	TCCR1C_FOC1C = 0x20 // Force Output Compare 1C

	// TCNT1L: Timer/Counter1 Bytes

	// TCNT1H: Timer/Counter1 Bytes
	TCNT1_TCNT1 = 0xffff // Timer/Counter1 bits

	// OCR1AL: Timer/Counter1 Output Compare Register A Bytes

	// OCR1AH: Timer/Counter1 Output Compare Register A Bytes
	OCR1A_OCR1A = 0xffff // Timer/Counter1 Output Compare A bits

	// OCR1BL: Timer/Counter1 Output Compare Register B Bytes

	// OCR1BH: Timer/Counter1 Output Compare Register B Bytes
	OCR1B_OCR1B = 0xffff // Timer/Counter1 Output Compare B bits

	// OCR1CL: Timer/Counter1 Output Compare Register C Bytes

	// OCR1CH: Timer/Counter1 Output Compare Register C Bytes
	OCR1C_OCR1C = 0xffff // Timer/Counter1 Output Compare C bits

	// ICR1L: Timer/Counter1 Input Capture Register Bytes

	// ICR1H: Timer/Counter1 Input Capture Register Bytes
	ICR1_ICR1 = 0xffff // Timer/Counter1 Input Capture bits

	// TIMSK1: Timer/Counter1 Interrupt Mask Register
	TIMSK1_ICIE1  = 0x20 // Timer/Counter1 Input Capture Interrupt Enable
	TIMSK1_OCIE1C = 0x8  // Timer/Counter1 Output Compare C Match Interrupt Enable
	TIMSK1_OCIE1B = 0x4  // Timer/Counter1 Output Compare B Match Interrupt Enable
	TIMSK1_OCIE1A = 0x2  // Timer/Counter1 Output Compare A Match Interrupt Enable
	TIMSK1_TOIE1  = 0x1  // Timer/Counter1 Overflow Interrupt Enable

	// TIFR1: Timer/Counter1 Interrupt Flag register
	TIFR1_ICF1  = 0x20 // Input Capture Flag 1
	TIFR1_OCF1C = 0x8  // Output Compare Flag 1C
	TIFR1_OCF1B = 0x4  // Output Compare Flag 1B
	TIFR1_OCF1A = 0x2  // Output Compare Flag 1A
	TIFR1_TOV1  = 0x1  // Timer/Counter1 Overflow Flag
)

// Bitfields for EXINT: External Interrupts
const (
	// EICRA: External Interrupt Control Register A
	EICRA_ISC3 = 0xc0 // External Interrupt Sense Control Bit
	EICRA_ISC2 = 0x30 // External Interrupt Sense Control Bit
	EICRA_ISC1 = 0xc  // External Interrupt Sense Control Bit
	EICRA_ISC0 = 0x3  // External Interrupt Sense Control Bit

	// EICRB: External Interrupt Control Register B
	EICRB_ISC7 = 0xc0 // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC6 = 0x30 // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC5 = 0xc  // External Interrupt 7-4 Sense Control Bit
	EICRB_ISC4 = 0x3  // External Interrupt 7-4 Sense Control Bit

	// EIMSK: External Interrupt Mask Register
	EIMSK_INT = 0xff // External Interrupt Request 7 Enable

	// EIFR: External Interrupt Flag Register
	EIFR_INTF = 0xff // External Interrupt Flags

	// PCMSK0: Pin Change Mask Register 0
	PCMSK0_PCINT = 0xff // Pin Change Mask 0

	// PCIFR: Pin Change Interrupt Flag Register
	PCIFR_PCIF0 = 0x1 // Pin Change Interrupt Flag 0

	// PCICR: Pin Change Interrupt Control Register
	PCICR_PCIE0 = 0x1 // Pin Change Interrupt Enable 0
)

// Bitfields for TWI: Two Wire Serial Interface
const (
	// TWAMR: TWI (Slave) Address Mask Register
	TWAMR_TWAM = 0xfe

	// TWCR: TWI Control Register
	TWCR_TWINT = 0x80 // TWI Interrupt Flag
	TWCR_TWEA  = 0x40 // TWI Enable Acknowledge Bit
	TWCR_TWSTA = 0x20 // TWI Start Condition Bit
	TWCR_TWSTO = 0x10 // TWI Stop Condition Bit
	TWCR_TWWC  = 0x8  // TWI Write Collition Flag
	TWCR_TWEN  = 0x4  // TWI Enable Bit
	TWCR_TWIE  = 0x1  // TWI Interrupt Enable

	// TWSR: TWI Status Register
	TWSR_TWS  = 0xf8 // TWI Status
	TWSR_TWPS = 0x3  // TWI Prescaler

	// TWAR: TWI (Slave) Address register
	TWAR_TWA   = 0xfe // TWI (Slave) Address register Bits
	TWAR_TWGCE = 0x1  // TWI General Call Recognition Enable Bit
)

// Bitfields for ADC: Analog-to-Digital Converter
const (
	// ADMUX: The ADC multiplexer Selection Register
	ADMUX_REFS  = 0xc0 // Reference Selection Bits
	ADMUX_ADLAR = 0x20 // Left Adjust Result
	ADMUX_MUX   = 0x1f // Analog Channel and Gain Selection Bits

	// ADCSRA: The ADC Control and Status register
	ADCSRA_ADEN  = 0x80 // ADC Enable
	ADCSRA_ADSC  = 0x40 // ADC Start Conversion
	ADCSRA_ADATE = 0x20 // ADC Auto Trigger Enable
	ADCSRA_ADIF  = 0x10 // ADC Interrupt Flag
	ADCSRA_ADIE  = 0x8  // ADC Interrupt Enable
	ADCSRA_ADPS  = 0x7  // ADC  Prescaler Select Bits

	// ADCL: ADC Data Register  Bytes

	// ADCH: ADC Data Register  Bytes
	ADC_ADC = 0x3ff // ADC Data Bits

	// DIDR0: Digital Input Disable Register 1
	DIDR0_ADC7D = 0x80 // ADC7 Digital input Disable
	DIDR0_ADC6D = 0x40 // ADC6 Digital input Disable
	DIDR0_ADC5D = 0x20 // ADC5 Digital input Disable
	DIDR0_ADC4D = 0x10 // ADC4 Digital input Disable
	DIDR0_ADC3D = 0x8  // ADC3 Digital input Disable
	DIDR0_ADC2D = 0x4  // ADC2 Digital input Disable
	DIDR0_ADC1D = 0x2  // ADC1 Digital input Disable
	DIDR0_ADC0D = 0x1  // ADC0 Digital input Disable

	// DIDR2: Digital Input Disable Register 1
	DIDR2_ADC13D = 0x20 // ADC13 Digital input Disable
	DIDR2_ADC12D = 0x10 // ADC12 Digital input Disable
	DIDR2_ADC11D = 0x8  // ADC11 Digital input Disable
	DIDR2_ADC10D = 0x4  // ADC10 Digital input Disable
	DIDR2_ADC9D  = 0x2  // ADC9 Digital input Disable
	DIDR2_ADC8D  = 0x1  // ADC8 Digital input Disable
)

// Bitfields for AC: Analog Comparator
const (
	// ACSR: Analog Comparator Control And Status Register
	ACSR_ACD  = 0x80 // Analog Comparator Disable
	ACSR_ACBG = 0x40 // Analog Comparator Bandgap Select
	ACSR_ACO  = 0x20 // Analog Compare Output
	ACSR_ACI  = 0x10 // Analog Comparator Interrupt Flag
	ACSR_ACIE = 0x8  // Analog Comparator Interrupt Enable
	ACSR_ACIC = 0x4  // Analog Comparator Input Capture Enable
	ACSR_ACIS = 0x3  // Analog Comparator Interrupt Mode Select bits

	// DIDR1
	DIDR1_AIN1D = 0x2 // AIN1 Digital Input Disable
	DIDR1_AIN0D = 0x1 // AIN0 Digital Input Disable
)

// Bitfields for CPU: CPU Registers
const (
	// SREG: Status Register
	SREG_I = 0x80 // Global Interrupt Enable
	SREG_T = 0x40 // Bit Copy Storage
	SREG_H = 0x20 // Half Carry Flag
	SREG_S = 0x10 // Sign Bit
	SREG_V = 0x8  // Two's Complement Overflow Flag
	SREG_N = 0x4  // Negative Flag
	SREG_Z = 0x2  // Zero Flag
	SREG_C = 0x1  // Carry Flag

	// OSCCAL: Oscillator Calibration Value
	OSCCAL_OSCCAL = 0xff // Oscillator Calibration

	// RCCTRL: Oscillator Control Register
	RCCTRL_RCFREQ = 0x1

	// CLKPR
	CLKPR_CLKPCE = 0x80
	CLKPR_CLKPS  = 0xf

	// SMCR: Sleep Mode Control Register
	SMCR_SM = 0xe // Sleep Mode Select bits
	SMCR_SE = 0x1 // Sleep Enable

	// RAMPZ: Extended Z-pointer Register for ELPM/SPM
	RAMPZ_RAMPZ = 0x3 // Extended Z-Pointer Value

	// GPIOR2: General Purpose IO Register 2
	GPIOR2_GPIOR = 0xff // General Purpose IO Register 2 bis

	// GPIOR1: General Purpose IO Register 1
	GPIOR1_GPIOR = 0xff // General Purpose IO Register 1 bis

	// GPIOR0: General Purpose IO Register 0
	GPIOR0_GPIOR07 = 0x80 // General Purpose IO Register 0 bit 7
	GPIOR0_GPIOR06 = 0x40 // General Purpose IO Register 0 bit 6
	GPIOR0_GPIOR05 = 0x20 // General Purpose IO Register 0 bit 5
	GPIOR0_GPIOR04 = 0x10 // General Purpose IO Register 0 bit 4
	GPIOR0_GPIOR03 = 0x8  // General Purpose IO Register 0 bit 3
	GPIOR0_GPIOR02 = 0x4  // General Purpose IO Register 0 bit 2
	GPIOR0_GPIOR01 = 0x2  // General Purpose IO Register 0 bit 1
	GPIOR0_GPIOR00 = 0x1  // General Purpose IO Register 0 bit 0

	// PRR1: Power Reduction Register1
	PRR1_PRUSB    = 0x80 // Power Reduction USB
	PRR1_PRTIM4   = 0x10 // Power Reduction Timer/Counter4
	PRR1_PRTIM3   = 0x8  // Power Reduction Timer/Counter3
	PRR1_PRUSART1 = 0x1  // Power Reduction USART1

	// PRR0: Power Reduction Register0
	PRR0_PRTWI    = 0x80 // Power Reduction TWI
	PRR0_PRTIM2   = 0x40 // Power Reduction Timer/Counter2
	PRR0_PRTIM0   = 0x20 // Power Reduction Timer/Counter0
	PRR0_PRTIM1   = 0x8  // Power Reduction Timer/Counter1
	PRR0_PRSPI    = 0x4  // Power Reduction Serial Peripheral Interface
	PRR0_PRUSART0 = 0x2  // Power Reduction USART
	PRR0_PRADC    = 0x1  // Power Reduction ADC

	// CLKSTA
	CLKSTA_RCON  = 0x2
	CLKSTA_EXTON = 0x1

	// CLKSEL1
	CLKSEL1_RCCKSEL = 0xf0
	CLKSEL1_EXCKSEL = 0xf

	// CLKSEL0
	CLKSEL0_RCSUT = 0xc0
	CLKSEL0_EXSUT = 0x30
	CLKSEL0_RCE   = 0x8
	CLKSEL0_EXTE  = 0x4
	CLKSEL0_CLKS  = 0x1
)

// Bitfields for PLL: Phase Locked Loop
const (
	// PLLCSR: PLL Status and Control register
	PLLCSR_PINDIV = 0x10 // PLL prescaler Bit 2
	PLLCSR_PLLE   = 0x2  // PLL Enable Bit
	PLLCSR_PLOCK  = 0x1  // PLL Lock Status Bit

	// PLLFRQ: PLL Frequency Control Register
	PLLFRQ_PINMUX = 0x80
	PLLFRQ_PLLUSB = 0x40
	PLLFRQ_PLLTM  = 0x30
	PLLFRQ_PDIV   = 0xf
)

// Bitfields for USB_DEVICE: USB Device Registers
const (
	// UEINT
	UEINT_EPINT = 0x7f // Endpoint interrupt bits

	// UEBCHX
	UEBCHX_BYCT = 0x7 // Byte count bits

	// UEBCLX
	UEBCLX_BYCT = 0xff // Byte count bits

	// UEDATX
	UEDATX_DAT = 0xff

	// UEIENX
	UEIENX_FLERRE   = 0x80
	UEIENX_NAKINE   = 0x40
	UEIENX_NAKOUTE  = 0x10
	UEIENX_RXSTPE   = 0x8
	UEIENX_RXOUTE   = 0x4
	UEIENX_STALLEDE = 0x2
	UEIENX_TXINE    = 0x1

	// UESTA1X
	UESTA1X_CTRLDIR = 0x4
	UESTA1X_CURRBK  = 0x3

	// UESTA0X
	UESTA0X_CFGOK   = 0x80
	UESTA0X_OVERFI  = 0x40
	UESTA0X_UNDERFI = 0x20
	UESTA0X_DTSEQ   = 0xc
	UESTA0X_NBUSYBK = 0x3

	// UECFG1X
	UECFG1X_EPSIZE = 0x70
	UECFG1X_EPBK   = 0xc
	UECFG1X_ALLOC  = 0x2

	// UECFG0X
	UECFG0X_EPTYPE = 0xc0
	UECFG0X_EPDIR  = 0x1

	// UECONX
	UECONX_STALLRQ  = 0x20
	UECONX_STALLRQC = 0x10
	UECONX_RSTDT    = 0x8
	UECONX_EPEN     = 0x1

	// UERST
	UERST_EPRST = 0x7f

	// UENUM
	UENUM_UENUM = 0x7

	// UEINTX
	UEINTX_FIFOCON  = 0x80
	UEINTX_NAKINI   = 0x40
	UEINTX_RWAL     = 0x20
	UEINTX_NAKOUTI  = 0x10
	UEINTX_RXSTPI   = 0x8
	UEINTX_RXOUTI   = 0x4
	UEINTX_STALLEDI = 0x2
	UEINTX_TXINI    = 0x1

	// UDMFN
	UDMFN_FNCERR = 0x10

	// UDFNUML

	// UDFNUMH
	UDFNUM_FNUM = 0x7ff // Frame number value

	// UDADDR
	UDADDR_ADDEN = 0x80
	UDADDR_UADD  = 0x7f

	// UDIEN
	UDIEN_UPRSME  = 0x40
	UDIEN_EORSME  = 0x20
	UDIEN_WAKEUPE = 0x10
	UDIEN_EORSTE  = 0x8
	UDIEN_SOFE    = 0x4
	UDIEN_SUSPE   = 0x1

	// UDINT
	UDINT_UPRSMI  = 0x40
	UDINT_EORSMI  = 0x20
	UDINT_WAKEUPI = 0x10
	UDINT_EORSTI  = 0x8
	UDINT_SOFI    = 0x4
	UDINT_SUSPI   = 0x1

	// UDCON
	UDCON_LSM    = 0x4 // USB low speed mode
	UDCON_RSTCPU = 0x8
	UDCON_RMWKUP = 0x2
	UDCON_DETACH = 0x1

	// USBCON: USB General Control Register
	USBCON_USBE    = 0x80
	USBCON_FRZCLK  = 0x20
	USBCON_OTGPADE = 0x10
	USBCON_VBUSTE  = 0x1

	// USBINT
	USBINT_VBUSTI = 0x1

	// USBSTA
	USBSTA_SPEED = 0x8
	USBSTA_VBUS  = 0x1

	// UHWCON
	UHWCON_UVREGE = 0x1
)
