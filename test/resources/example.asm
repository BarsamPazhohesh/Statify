section .data
    ; This is a single-line comment
    message db 'Hello, World!', 0

section .text
    global _start

_start:
    ; This is another single-line comment
    ; The next few lines print the message to the screen

    mov eax, 4          ; syscall number for sys_write
    mov ebx, 1          ; file descriptor 1 (stdout)
    mov ecx, message    ; pointer to message
    mov edx, 13         ; length of message
    int 0x80            ; interrupt to make the syscall

    ; This is the final comment before the exit
    mov eax, 1          ; syscall number for sys_exit
    xor ebx, ebx        ; exit status 0
    int 0x80            ; interrupt to make the syscall

; This is a multi-line comment
; that spans several lines.
; It doesn't affect the program execution.
; These comments can be used for explanation or notes.

