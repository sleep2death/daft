<template>
  <div class="cols columns">
    <div class="column is-one-fifth is-mobile">
      <div class="menu">
        <p class="menu-label">
          Select {{ name }}
        </p>
        <ul class="menu-list">
          <router-link
            v-for="(item, index) in list"
            :key="index"
            tag="li"
            :to="{name: name, params:{id: item.id}}"
            exact
          >
            <a>{{ item.id + ' - ' + item.name }}</a>
          </router-link>
        </ul>
      </div>
    </div>
    <div class="column">
      <router-view />
    </div>
  </div>
</template>

<script>
import { mapState } from 'vuex'

export default {
  name: 'List',
  props: {
    name: {
      type: String,
      default: function () {
        return 'LIST'
      }
    },
    list: {
      type: Array,
      default: function () {
        return null
      }
    }
  },
  computed: mapState(['actors', 'route', 'current']),
  updated: function () {
    if (this.route.params.id && this.route.params.id !== this.current) {
      this.$store.commit('setCurrent', Number(this.route.params.id))
    }
    if (!this.route.params.id && this.current > 0) {
      this.$router.push({ name: this.name, params: { id: this.current } })
    }
  }
}
</script>

<style scoped>
  .cols {
    margin-top:  0.5rem;
  }
</style>
