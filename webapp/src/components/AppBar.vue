<template>
  <nav
    class="navbar is-primary"
    role="navigation"
    aria-label="main navigation"
  >
    <div class="navbar-brand">
      <div class="navbar-item">
        <span class="icon">
          <i class="far fa-rabbit" />
        </span>
        <span class="logo">Rabbits Maker</span>
      </div>
      <a
        role="button"
        class="navbar-burger burger"
        aria-label="menu"
        aria-expanded="false"
        data-target="n-menu"
      >
        <span aria-hidden="true" />
        <span aria-hidden="true" />
        <span aria-hidden="true" />
      </a>
    </div>

    <div
      id="n-menu"
      class="navbar-menu"
    >
      <div class="navbar-start">
        <router-link
          to="/"
          tag="a"
          class="navbar-item"
          exact
        >
          主页
        </router-link>
        <router-link
          tag="a"
          to="/database"
          class="navbar-item"
        >
          数据库
        </router-link>
      </div>
    </div>
  </nav>
</template>

<script>

export default {
  name: 'AppBar',
  mounted: function () {
    // Get all "navbar-burger" elements
    const $navbarBurgers = Array.prototype.slice.call(document.querySelectorAll('.navbar-burger'), 0)
    let $target = null

    // Check if there are any navbar burgers
    if ($navbarBurgers.length > 0) {
      // Add a click event on each of them
      $navbarBurgers.forEach(el => {
        el.addEventListener('click', () => {
          // Get the target from the "data-target" attribute
          const target = el.dataset.target
          $target = document.getElementById(target)
          // Toggle the "is-active" class on both the "navbar-burger" and the "navbar-menu"
          el.classList.toggle('is-active')
          $target.classList.toggle('is-active')
        })
      })
    }
  }
}
</script>
