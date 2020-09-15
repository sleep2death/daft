<template>
  <div>
    <tabs :data="tabs" />
    <router-view />
  </div>
</template>

<script>
import Tabs from '@/components/Tabs.vue'
import { mapState } from 'vuex'

export default {
  name: 'Database',
  components: {
    Tabs
  },
  computed: mapState({
    tabs: state => state.dbTabs
  })
}
</script>
