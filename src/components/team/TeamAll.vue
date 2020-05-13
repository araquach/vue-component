<template>
    <div>
        <h1 class="title is-3">Meet the full team</h1>
        <div class="columns is-multiline">
            <div v-for="(tm, id) in teamMembers" class="column is-4">
                <router-link :to="{ name: 'team-detail', params: { slug: tm.slug }}" :tm="tm">
                    <div class="box">
                        <h2 class="title is-6">{{ tm.name }}</h2>
                        <p>{{ tm.para_1 }}</p>
                    </div>
                </router-link>
            </div>
        </div>
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
                'teamMembers'
            ])
        }
    }
</script>