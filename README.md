# Proof of Concept

> [!WARNING]  
> **Abandoned – But Not in Vain!**
> 
> This was a great educational exercise. While I've learned static-php & phpmicro already solve the problem, building cross-platform executables is still complex.
> 
> I'm building a opinionated CLI tool to make it stupidly simple — stay tuned! 👀

Package a PHP script as a standalone cross-platform binary

**Note this PoC only compiles for Mac arm at this time.**

1. This project uses [NativePHP's static binaries](https://github.com/NativePHP/php-bin), so after cloning first run `composer install`.
2. You need to have Go installed on your machine to cross-compile for your desired architecture.

To compile the script in `php/app.php` run:

```bash
./build.sh
```

Afterwards a binary should be present in `./build/mac-arm` 🚀
