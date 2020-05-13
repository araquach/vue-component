const state = {
    teamMembers: []
}

const getters = {
    teamMembers: state => state.teamMembers,
    teamMember: state => slug => {
        return state.teamMembers.find(tm => tm.slug === slug)
    }
}

const mutations = {
    SET_TEAM_MEMBERS (state, teamMembers) {
        state.teamMembers = teamMembers
    }
}

const actions = {
    loadTeamMembers ({ commit }) {
        axios
            .get('/api/team')
            .then(r => r.data)
            .then(teamMembers => {
                commit('SET_TEAM_MEMBERS', teamMembers)
            })
    },
}

export default {
    state,
    getters,
    mutations,
    actions
}