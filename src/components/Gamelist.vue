<template>
    <div class="section">
        <h1>Games</h1>
        <ul>
            <li v-for="game in games" :key="game.id">
                <router-link :to="{ name: 'game-rules', params: { id: game.id } }">
                    {{ game.title }}
                </router-link>
            </li>
        </ul>
    </div>
</template>

<script>
    import games from "../data/games"

    export default {
        data() {
            return {
                games: games
            }
        }
    }
</script>