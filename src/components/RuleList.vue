<template>
    <div>
        <h3>Rules</h3>
        <ul>
            <li v-for="rule in game.rules" :key="rule.id">
                <router-link :to="{ name: 'Game Rule', params: { ruleId: rule.id } }">
                    {{ excerpt(rule.text) }}
                </router-link>
            </li>
        </ul>
    </div>
</template>

<script>
    import games from "../data/games";

    export default {
        name: "rule-list",
        data() {
            return {
                game: games.filter(game => game.id === parseInt(this.$route.params.id))[0]
            };
        },
        methods: {
            excerpt(text) {
                return text.substr(0, 100) + "...";
            }
        }
    };
</script>
