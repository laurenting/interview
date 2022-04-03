<template>
  <div style="width: 960px;margin: 10px auto">
    <el-form
        label-position="top"
        label-width="950px"
        :model="list"
        style="max-width: 960px"
    >
      <el-form-item
          v-for="(v,i) in list"
          :key="i"
          :label="`Q${i+1}: ${v.question}`">
        <el-input :disabled="!editable" v-model="v.answer"/>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: "Questionnaire",
  props: {
    questionList: {
      type: Array,
      default: () => {
        return []
      }
    },
    editable: {
      type: Boolean,
      default: true
    }
  },
  watch: {
    questionList: {
      handler(val, oldVal)  {
        this.initList(val)
      },
      deep: true
    }
    },
  created() {
    this.initList(this.questionList)
  },
  data() {
    return {
      list: []
    }
  },
  methods: {
    initList(val) {
      this.list = val.map(q => {
        return {
          question: q.value,
          answer: q.answer ?? '',
          timestamp: q.key ?? 0,
        }
      }).sort((x,y) => y.timestamp - x.timestamp)
    }
  }
}
</script>


<style scoped>

</style>