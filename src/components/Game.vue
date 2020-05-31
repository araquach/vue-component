<template>
    <div class="section">
        <h1>{{game.title}}</h1>

        <router-view name="page"/>

        <modal v-if="showModal" ref="modal">
            <router-view name="rule"/>
        </modal>
    </div>
</template>

<script>
    import games from '../data/games'
    import Modal from './Modal'

    export default {
        name: 'Game',
        components: {
            Modal
        },

        data() {
            return {
                game: games.filter(
                    game => game.id === parseInt(this.$route.params.id)
                )[0],
                showModal: this.$route.meta.showModal
            }
        },

        watch: {
            '$route.meta'({ showModal }) {
                this.showModal = showModal
            }
        }
    }
</script>