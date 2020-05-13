<template>

    <div>
        <h1 class="title is-3">Team Detail View</h1>
<!--        <h2 class="title is-4">{{ tm.name }}</h2>-->
        <p>{{ tm.para_1 }}</p>
        <p>{{ tm.para_2 }}</p>
    </div>
</template>

<script>
    import {mapGetters} from 'vuex'

    export default {
        mounted() {
            this.$store.dispatch('loadTeamMembers')
        },

        computed: {
            ...mapGetters([
                'teamMember'
            ]),

            tm() {
                return this.teamMember(this.$route.params.slug)
            }
        }
    }
</script>