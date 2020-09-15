<template>
  <div class="section columns">
    <div class="column is-one-fifth is-mobile">
      <div class="menu">
        <p class="menu-label">
          Actors
        </p>
        <ul class="menu-list">
          <router-link
            v-for="(actor, index) in actors"
            :key="index"
            tag="li"
            :to="'/database/actors/' + actor.id"
            exact
          >
            <a>{{ actor.name }}</a>
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
  name: 'Actors',
  computed: mapState(['actors']),
  mounted: function () {
    this.$store.dispatch('getActors')
  }
}
</script>
