<template>
  <div>
    <div class="tabs">
      <ul>
        <router-link
          v-for="(item, index) in data"
          :key="index"
          tag="li"
          :to="item.link"
        >
          <a>
            <span class="icon">
              <i :class="item.icon" />
            </span>
            <span>{{ item.label }}</span>
          </a>
        </router-link>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: 'Tabs',
  components: {
  },
  props: {
    data: {
      type: Array,
      default: function () {
        return ['Actors']
      }
    }
  }
}
</script>
