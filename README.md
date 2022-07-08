<div id="top"></div>
<!--
*** Thanks for checking out the Best-README-Template. If you have a suggestion
*** that would make this better, please fork the repo and create a pull request
*** or simply open an issue with the tag "enhancement".
*** Don't forget to give the project a star!
*** Thanks again! Now go create something AMAZING! :D
-->



<!-- PROJECT SHIELDS -->
<!--
*** I'm using markdown "reference style" links for readability.
*** Reference links are enclosed in brackets [ ] instead of parentheses ( ).
*** See the bottom of this document for the declaration of the reference variables
*** for contributors-url, forks-url, etc. This is an optional, concise syntax you may use.
*** https://www.markdownguide.org/basic-syntax/#reference-style-links
-->

[![LinkedIn][linkedin-shield]][linkedin-url]



<!-- PROJECT LOGO -->
<br />
<div align="center">

  <h3 align="center">Golang JWT CRUD</h3>

  <p align="center">
    For detail API Documentation please check bellow <br>
    https://documenter.getpostman.com/view/9974746/UzJLPGhB
  </p>
</div>



<!-- GETTING STARTED -->
## Getting Started

Before you running this on your computer make sure you install docker and docker comopse

### Prerequisites

To install docker and docker compose https://www.docker.com/


### Installation

_Here is how to running project from docker._

1. First move to project folder
2. Running the docker-compose-local.yml
   ```sh
   docker-compose -f docker-compose-local.yml up -d
   ```
3. Running query on database to create table and init user
> CREATE TABLE users (
    id CHAR(36) PRIMARY KEY,
    username VARCHAR(30) NOT NULL, 
    password VARCHAR(255) NOT NULL, 
    firstname VARCHAR(30) NOT NULL,
    lastname VARCHAR(30) NOT NULL,
    email VARCHAR(50) NOT NULL,
    phone VARCHAR(30) NOT NULL,
    user_role VARCHAR(30) NOT NULL,
    profile_picture varchar(255) NOT NULL,
    is_deleted int default 0,
    refresh_token varchar(255),
    reset_password varchar(255),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
); 

> INSERT INTO `users` (`id`, `username`, `password`, `firstname`, `lastname`, `email`, `phone`, `user_role`, `profile_picture`, `is_deleted`, `refresh_token`, `reset_password`, `created_at`, `updated_at`)
VALUES
	('3b2c1306-12d5-4357-aab8-bbf49c9a6a20', 'user', '$2a$14$kZgfh5ttpvw9YfqSx496UO8FGLd08WQ/NWZtXFweuIhZ5A9KPeprC', 'angga', 'pamungkas', 'kap21kapuser@gmail.com', '3434534', 'user', '576113_realsam_o2_logo_partner.png', 0, NULL, NULL, '2022-07-07 15:54:52', '2022-07-07 15:54:52'),
	('fc7f19bd-a9a8-445c-bf5f-6e73936ff24f', 'admin', '$2a$14$YcBNkbqjbzUzEH3EQzPIA.0U.KXzFOFn0xAc6vgNhAvygbjmCxrhq', 'angga', 'pamungkas', 'kap21kap@gmail.com', '3434534', 'admin', '831252_realsam_o2_logo_partner.png', 0, NULL, NULL, '2022-07-07 15:54:24', '2022-07-07 15:54:24');


4. Now you can access the project from localhost:9090

<p align="right">(<a href="#top">back to top</a>)</p>




### Usage API

_You can use the API from local or server (I have been installed)._

API END POINT SERVER : http://103.31.38.61:9090/ <br>
API END POINT LOCAL (install from docker) : http://localhost:9090/<br>
<br>
Credentials<br>
<br>
Admin<br>
user : admin<br>
password : admin<br>
<br>
User<br>
user : user<br>
password : user
<br>


<p align="right">(<a href="#top">back to top</a>)</p>


<!-- CONTACT -->
## Contact

Krisna Angga Pamungkas - krisnaanggapamungkas@gmail.com

Project Link: [https://github.com/kapamungkas](https://github.com/kapamungkas)

<p align="right">(<a href="#top">back to top</a>)</p>



<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/othneildrew/Best-README-Template.svg?style=for-the-badge
[contributors-url]: https://github.com/othneildrew/Best-README-Template/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/othneildrew/Best-README-Template.svg?style=for-the-badge
[forks-url]: https://github.com/othneildrew/Best-README-Template/network/members
[stars-shield]: https://img.shields.io/github/stars/othneildrew/Best-README-Template.svg?style=for-the-badge
[stars-url]: https://github.com/othneildrew/Best-README-Template/stargazers
[issues-shield]: https://img.shields.io/github/issues/othneildrew/Best-README-Template.svg?style=for-the-badge
[issues-url]: https://github.com/othneildrew/Best-README-Template/issues
[license-shield]: https://img.shields.io/github/license/othneildrew/Best-README-Template.svg?style=for-the-badge
[license-url]: https://github.com/othneildrew/Best-README-Template/blob/master/LICENSE.txt
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/kapamungkas/
[product-screenshot]: images/screenshot.png
[Next.js]: https://img.shields.io/badge/next.js-000000?style=for-the-badge&logo=nextdotjs&logoColor=white
[Next-url]: https://nextjs.org/
[React.js]: https://img.shields.io/badge/React-20232A?style=for-the-badge&logo=react&logoColor=61DAFB
[React-url]: https://reactjs.org/
[Vue.js]: https://img.shields.io/badge/Vue.js-35495E?style=for-the-badge&logo=vuedotjs&logoColor=4FC08D
[Vue-url]: https://vuejs.org/
[Angular.io]: https://img.shields.io/badge/Angular-DD0031?style=for-the-badge&logo=angular&logoColor=white
[Angular-url]: https://angular.io/
[Svelte.dev]: https://img.shields.io/badge/Svelte-4A4A55?style=for-the-badge&logo=svelte&logoColor=FF3E00
[Svelte-url]: https://svelte.dev/
[Laravel.com]: https://img.shields.io/badge/Laravel-FF2D20?style=for-the-badge&logo=laravel&logoColor=white
[Laravel-url]: https://laravel.com
[Bootstrap.com]: https://img.shields.io/badge/Bootstrap-563D7C?style=for-the-badge&logo=bootstrap&logoColor=white
[Bootstrap-url]: https://getbootstrap.com
[JQuery.com]: https://img.shields.io/badge/jQuery-0769AD?style=for-the-badge&logo=jquery&logoColor=white
[JQuery-url]: https://jquery.com 
