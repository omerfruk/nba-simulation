<template>
  <div>
    <b-navbar type="dark" variant="dark">
      <b-navbar-nav align="center">
        <b-nav-item class="" href="#">NBA Simulator</b-nav-item>
      </b-navbar-nav>
    </b-navbar>

    <b-container>
      <b-row v-for="(val,index) in value" :key="index" class="m-3">
        <b-col>
          <b-card
              :header="val.MatchName"
              border-variant="primary"
              header-bg-variant="primary"
              header-text-variant="white"
              align="center"
          >
            <div v-for="(data, index2) in val.TeamScoreVM" :key="index2">
              <h3>{{ data.TeamName }}</h3>
              <table class="table mb-3">
                <thead>
                <tr>
                  <th scope="col">Player Id</th>
                  <th scope="col">Player Name</th>
                  <th scope="col">Shot Point</th>
                  <th scope="col">Player Status</th>
                </tr>
                </thead>
                <tbody>
                <tr v-for="(player , ind) in  data.PlayerScoreVM" :key="ind">
                  <th scope="row">{{ player.PlayerId }}</th>
                  <th scope="row"> {{ player.PlayerName }}</th>
                  <th>{{ player.PlayerScoreType == 0 ? "Assist" : player.PlayerScoreType + " Point" }}</th>
                  <th>{{ player.PlayerType == 1 ? "Player" : "Substitute"}}</th>
                </tr>
                </tbody>
              </table>
            </div>
          </b-card>
        </b-col>
      </b-row>

    </b-container>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "home",
  data() {
    return {
      value: {},
      items: [
        {age: 40, first_name: 'Dickerson', last_name: 'Macdonald'},
        {age: 21, first_name: 'Larsen', last_name: 'Shaw'},
        {age: 89, first_name: 'Geneva', last_name: 'Wilson'},
        {age: 38, first_name: 'Jami', last_name: 'Carney'}
      ]
    }
  },
  created() {
    this.getMatch()
  },
  methods: {
    async getMatch() {
      await axios.get('http://localhost:3000').then(response => (this.value = response.data))
    }
  },
}
</script>

<style scoped>

</style>