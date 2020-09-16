<template>
  <list
    name="actors"
    :list="actors"
  />
</template>

<script>
import { mapState } from 'vuex'
import List from '@/components/List'

export default {
  name: 'Actors',
  components: {
    List
  },
  computed: mapState(['actors']),
  mounted: function () {
    this.$store.dispatch('getActors')
  }
}
</script>

<style scoped>
  .cols {
    margin-top:  0.5rem;
  }
</style>
