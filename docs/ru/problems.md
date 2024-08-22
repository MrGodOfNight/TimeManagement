# Проблемы

## Сервер

### При запуске сервера сообщения в консоли не цветные на windows.

Причина: проблема с отключенной поддержкой ANSI в консоли windows.

Решение: включите поддержку ANSI в консоли windows (напишите в powershell `Set-ItemProperty -Path "HKCU:\Console" -Name VirtualTerminalLevel -Value 1` Работает только на windows 10 и выше).
