<template>
    <div>
        <p>{{ rule.text }}</p>

        <ul>
            <li>
                <button>Delete</button>
            </li>
        </ul>
    </div>
</template>

<script>
    import games from "../data/games";
    export default {
        name: "rule",
        computed: {
            rule() {
                const { id, ruleId } = this.$route.params;
                const game = games.find(game => game.id === parseInt(id));
                return game.rules.find(rule => rule.id === parseInt(ruleId));
            }
        }
    };
</script>