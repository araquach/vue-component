<template>
    <div>
        <p class="has-text-white">Name: {{ fullName }}</p>
        <p class="has-text-white">Age: {{ person.age }}</p>
    </div>
</template>

<script>
    import people from "../data/people"
    export default {
        data() {
            return {
                person: people.filter(
                    person => person.id === parseInt(this.$route.params.id))[0]
            }
        },

        computed: {
            fullName() {
                return `${this.person.firstName} ${this.person.lastName}`
            }
        }
    }
</script>


