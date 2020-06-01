<template>
    <div>
        <p>{{ rule.text }} Adam Carter</p>
    </div>
</template>

<script>
    import games from "../data/games"
    export default {
        computed: {
            rule() {
                const { id, ruleId } = this.$route.params;
                const game = games.find(game => game.id === parseInt(id));
                return game.rules.find(rule => rule.id === parseInt(ruleId));
            }
        }
    };
</script>