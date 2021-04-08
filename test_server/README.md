# Remote test server🖥

## Params

 * **IP**: `185.253.218.139`
 * **domain**: `fedorenka.online`
 * **OS**: *Debian* 10

## Subdomains

 * `apple`: [Apple🥧pie](https://github.com/FedorenkaAvenue/Apple_pie)
 * `beria`: [Beria🦅](https://github.com/FedorenkaAvenue/Beria)
 * `docs`: [Magallanes⛵️](https://github.com/FedorenkaAvenue/Magallanes)

## TCP/IP Network

 * [website📬](../website)

    * `80`

 * [Apple🥧pie](https://github.com/FedorenkaAvenue/Apple_pie)

    * `81`: [Nginx](https://github.com/FedorenkaAvenue/Apple_pie/tree/master/nginx) entrypoint container
    * `6379`: [Redis](https://github.com/FedorenkaAvenue/Apple_pie/tree/master/redis)
    * `5432`: [PostgreSQL](https://github.com/FedorenkaAvenue/Apple_pie/tree/master/postgres)

 * [Beria🦅](https://github.com/FedorenkaAvenue/Beria)
 
    * `82`: [Nginx](https://github.com/FedorenkaAvenue/Beria/tree/master/nginx) entrypoint container
    * `6380`: [Redis](https://github.com/FedorenkaAvenue/Beria/tree/master/redis)

 * [Magallanes⛵️](https://github.com/FedorenkaAvenue/Magallanes)
    * `911`: [docker](https://github.com/FedorenkaAvenue/Magallanes) container

## Auth

 * user: `root`
 * pass: `fedorenka`

## Guides

<details>
   <summary>📔генерирование <i>SSL</i></summary>
      SSl сгенерирован через <a href="https://certbot.eff.org/">Certbot</a>.<br>
      Местонахождение сертификата: <code>/etc/letsencrypt/live/fedorenka.online/</code><br>
      Создать сертификат: <code>certbot --nginx -d sub.domain.com</code><br>
</details>  

<details>
   <summary>📔<i>Nginx</i> auth</summary>
      <code>apache2-utils</code> для использования <code>htpasswd</code><br>
      <code>htpasswd -c /etc/nginx/.http_auth USER</code> для генерирования юзера и пароля
</details>
