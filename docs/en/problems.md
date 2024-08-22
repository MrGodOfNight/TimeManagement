# Problems

## Server

### When the server starts, the messages in the console are not colored on windows.

Reason: the problem is with the support of ANSI in the windows console.

Solution: enable the support of ANSI in the windows console (write in powershell `Set-ItemProperty -Path "HKCU:\Console" -Name VirtualTerminalLevel -Value 1` Works only on windows 10 and above).