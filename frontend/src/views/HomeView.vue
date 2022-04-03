<template>
  <div class="app-container">
    <div class="top">
      <h1>
        Please Choose Your Identity
      </h1>
    </div>
    <div class="body">
      <div class="btn-box">
        <div>
          <el-card class="box-card" shadow="hover">
            <template #header>
              <div class="card-header">
                <h4>I'm Interviewer
                  <el-icon :size="20">
                    <user-filled/>
                  </el-icon>
                </h4>
              </div>
            </template>
            <div>
              <div class="id-desc-info">
                <el-link :underline="false" @click="$router.push('/interviewer/create_questionnaire')">
                  <el-icon>
                    <document-add/>
                  </el-icon>
                  Create questionnaire
                </el-link>
              </div>
              <div class="id-desc-info">
                <el-link :underline="false" @click="$router.push('/interviewer/view_questionnaire')">
                  <el-icon>
                    <document-checked/>
                  </el-icon>
                  Open completed questionnaire
                </el-link>
              </div>
              <div class="id-desc-info">...</div>
            </div>
          </el-card>
        </div>
        <div @click="goInterviewee">
          <el-card class="box-card" shadow="hover">
            <template #header>
              <div class="card-header">
                <h4>I'm Interviewee
                  <el-icon :size="20">
                    <user/>
                  </el-icon>
                </h4>
              </div>
            </template>
            <div>
              <div class="id-desc-info">
                <el-link :underline="false">
                  <el-icon>
                    <mic/>
                  </el-icon>
                  Fill out questionnaire
                </el-link>
              </div>
              <div class="id-desc-info">...</div>
            </div>
          </el-card>
        </div>
      </div>
    </div>
    <div class="footer">
      Copyright © 2022 Renting Liu and DAOZHUODA. All rights reserved.
    </div>
  </div>
</template>

<script>

export default {
  name: 'HomeView',
  components: {},
  methods: {
    goInterviewee() {
      this.$router.push('/')
    }
  }
}
</script>

<style scoped lang="scss">
.app-container {
  display: flex;
  flex-direction: column;
}

.top {
  text-align: center;
  padding: 20px 0;

  .h3 {
    font-size: 50px;
  }
}

.body {
  flex: 1;
}

.body .btn-box {
  display: flex;
  flex-direction: row;
  flex-wrap: nowrap;
  justify-content: space-around;
}

.box-card {
  height: 525px;
  width: 380px;
  border-radius: 5px;
  //cursor: pointer;
}

.footer {
  text-align: center;
  padding: 10px 0
}

.id-desc-info {
  text-align: center;
  margin: 20px 0;
}
</style>
