
import { ref } from "vue"
import { defineStore } from "pinia"

import { getCategories } from "@/services/category"
import type { Category } from "@/types/category"

export const useCategoryStore = defineStore("category", () => {
  const categories = ref<Category[]>([])
  const isLoading = ref(false)
  const isLoaded = ref(false)

  async function loadCategories(force = false) {
    if (isLoaded.value && !force) {
      return
    }

    try {
      isLoading.value = true

      categories.value = await getCategories()
      isLoaded.value = true
    } finally {
      isLoading.value = false
    }
  }

  function addCategory(category: Category) {
    categories.value.push(category)
  }

  function removeCategory(categoryId: number) {
    categories.value = categories.value.filter(
      (category) => category.id !== categoryId
    )
  }

  return {
    categories,
    isLoading,
    isLoaded,
    loadCategories,
    addCategory,
    removeCategory,
  }
})