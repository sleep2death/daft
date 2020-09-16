import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    tabs: ['Actors', 'Classes', 'Items', 'Skills', 'State', 'Weapons', 'Armores', 'Enemies'],
    dbTabs: [
      { label: '角色', link: '/database/actors', icon: 'far fa-user' },
      { label: '职业', link: '/database/classes', icon: 'far fa-swords' },
      { label: '技能', link: '/database/skills', icon: 'far fa-book-spells' }
    ],
    actors: [],
    current: 0
  },
  getters: {
    getActorByID: (state) => (id) => {
      return state.actors.find(actor => actor.id === id)
    }
  },
  mutations: {
    setActors (state, payload) {
      // update actors
      state.actors = payload
      if (state.current === 0) {
        state.current = state.actors[0].id
      }
    },
    setCurrent (state, payload) {
      state.current = payload
    }
  },
  actions: {
    getActors ({ commit }) {
      commit('setActors', [
        { name: 'aspirin', id: 1 },
        { name: 'shi', id: 2 }
      ])
    }
  },
  modules: {
  }
})
