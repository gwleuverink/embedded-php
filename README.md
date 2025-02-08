# Proof of Concept

Package a PHP script as a standalone cross-platform binary

**Note this PoC only compiles for Mac at this time.**

1. This project uses [NativePHP's static php binaries](https://github.com/NativePHP/php-bin), so after cloning first runs `composer install`.
2. You need to have Go installed on your machine to cross-compile for your desired architecture.

To compile the script in `php/app.php` run:

```bash
./build.sh
```
