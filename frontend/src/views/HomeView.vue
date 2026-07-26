<script setup lang="ts">
import { ref, computed } from 'vue'

const today = new Date()
const currentYear = ref(today.getFullYear())
const currentMonth = ref(today.getMonth())

const selectedDate = ref<number | null>(today.getDate())
const showAddModal = ref(false)
const newEvent = ref({ title: '', amount: '', type: 'expense' as 'income' | 'expense', category: '' })

const thaiMonths = [
  'มกราคม','กุมภาพันธ์','มีนาคม','เมษายน','พฤษภาคม','มิถุนายน',
  'กรกฎาคม','สิงหาคม','กันยายน','ตุลาคม','พฤศจิกายน','ธันวาคม',
]
const thaiDays = ['อา', 'จ', 'อ', 'พ', 'พฤ', 'ศ', 'ส']

const monthLabel = computed(() => `${thaiMonths[currentMonth.value]} ${currentYear.value + 543}`)
const daysInMonth = computed(() => new Date(currentYear.value, currentMonth.value + 1, 0).getDate())
const firstDayOfMonth = computed(() => new Date(currentYear.value, currentMonth.value, 1).getDay())

const events = ref<Record<string, { title: string; amount: number; type: 'income' | 'expense'; category: string }[]>>({
  [`${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2,'0')}-01`]: [
    { title: 'เช่าบ้าน', amount: 8500, type: 'expense', category: 'ที่พัก' },
  ],
  [`${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2,'0')}-05`]: [
    { title: 'เงินเดือน', amount: 45000, type: 'income', category: 'รายได้' },
    { title: 'ค่าอาหาร', amount: 350, type: 'expense', category: 'อาหาร' },
  ],
  [`${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2,'0')}-10`]: [
    { title: 'ค่าไฟ', amount: 1200, type: 'expense', category: 'สาธารณูปโภค' },
    { title: 'ค่าน้ำ', amount: 320, type: 'expense', category: 'สาธารณูปโภค' },
  ],
  [`${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2,'0')}-15`]: [
    { title: 'รายได้เสริม', amount: 5000, type: 'income', category: 'รายได้' },
  ],
  [`${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2,'0')}-${String(today.getDate()).padStart(2,'0')}`]: [
    { title: 'กาแฟ & อาหารเที่ยง', amount: 280, type: 'expense', category: 'อาหาร' },
  ],
  [`${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2,'0')}-20`]: [
    { title: 'ค่าเน็ต + มือถือ', amount: 799, type: 'expense', category: 'สื่อสาร' },
  ],
  [`${today.getFullYear()}-${String(today.getMonth() + 1).padStart(2,'0')}-25`]: [
    { title: 'ออมทรัพย์', amount: 10000, type: 'expense', category: 'ออม' },
  ],
})

function getDateKey(day: number) {
  return `${currentYear.value}-${String(currentMonth.value + 1).padStart(2,'0')}-${String(day).padStart(2,'0')}`
}
function getEventsForDay(day: number) { return events.value[getDateKey(day)] || [] }
function hasIncome(day: number) { return getEventsForDay(day).some(e => e.type === 'income') }
function hasExpense(day: number) { return getEventsForDay(day).some(e => e.type === 'expense') }
function isToday(day: number) {
  return day === today.getDate() && currentMonth.value === today.getMonth() && currentYear.value === today.getFullYear()
}

function prevMonth() {
  if (currentMonth.value === 0) { currentMonth.value = 11; currentYear.value-- }
  else currentMonth.value--
  selectedDate.value = null
}
function nextMonth() {
  if (currentMonth.value === 11) { currentMonth.value = 0; currentYear.value++ }
  else currentMonth.value++
  selectedDate.value = null
}
function goToday() {
  currentYear.value = today.getFullYear()
  currentMonth.value = today.getMonth()
  selectedDate.value = today.getDate()
}

const selectedDayEvents = computed(() => selectedDate.value != null ? getEventsForDay(selectedDate.value) : [])
const selectedDayTotal = computed(() => {
  const inc = selectedDayEvents.value.filter(e => e.type === 'income').reduce((s, e) => s + e.amount, 0)
  const exp = selectedDayEvents.value.filter(e => e.type === 'expense').reduce((s, e) => s + e.amount, 0)
  return { income: inc, expense: exp }
})
const monthSummary = computed(() => {
  let income = 0, expense = 0
  const prefix = `${currentYear.value}-${String(currentMonth.value + 1).padStart(2,'0')}-`
  Object.entries(events.value).forEach(([key, evs]) => {
    if (key.startsWith(prefix)) evs.forEach(e => e.type === 'income' ? (income += e.amount) : (expense += e.amount))
  })
  return { income, expense, net: income - expense }
})

function formatMoney(n: number) { return n.toLocaleString('th-TH') }

function addEvent() {
  if (!newEvent.value.title || !newEvent.value.amount || selectedDate.value == null) return
  const key = getDateKey(selectedDate.value)
  if (!events.value[key]) events.value[key] = []
  events.value[key].push({
    title: newEvent.value.title,
    amount: Number(newEvent.value.amount),
    type: newEvent.value.type,
    category: newEvent.value.category || 'ทั่วไป',
  })
  newEvent.value = { title: '', amount: '', type: 'expense', category: '' }
  showAddModal.value = false
}

const calendarDays = computed(() => {
  const days: (number | null)[] = []
  for (let i = 0; i < firstDayOfMonth.value; i++) days.push(null)
  for (let d = 1; d <= daysInMonth.value; d++) days.push(d)
  return days
})
</script>

<template>
  <main class="flex-1 min-w-0 p-7 overflow-y-auto bg-[#0f1117] flex flex-col gap-6">

    <!-- ── Page Header ── -->
    <div class="flex items-start justify-between">
      <div>
        <h1 class="text-[26px] font-extrabold text-white tracking-tight">ปฏิทินการเงิน</h1>
        <p class="text-[13px] text-slate-500 mt-1">ติดตามรายรับ-รายจ่ายรายวัน</p>
      </div>
      <div class="flex gap-2.5 items-center flex-shrink-0">
        <button
          class="px-[18px] py-2 rounded-[9px] border border-white/10 bg-white/[0.06] text-slate-300 text-[13px] cursor-pointer transition-colors hover:bg-white/[0.11]"
          @click="goToday"
        >วันนี้</button>
        <button
          class="flex items-center gap-1.5 px-[18px] py-2 rounded-[9px] border-none bg-gradient-to-br from-[#6c63ff] to-[#8b84ff] text-white text-[13px] font-semibold cursor-pointer shadow-[0_4px_14px_rgba(108,99,255,0.4)] transition-all hover:-translate-y-px hover:shadow-[0_6px_20px_rgba(108,99,255,0.5)] disabled:opacity-45 disabled:cursor-not-allowed disabled:translate-y-0 disabled:shadow-none"
          @click="showAddModal = true"
          :disabled="selectedDate == null"
        >
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2.5" width="15" height="15">
            <line x1="12" y1="5" x2="12" y2="19"/><line x1="5" y1="12" x2="19" y2="12"/>
          </svg>
          เพิ่มรายการ
        </button>
      </div>
    </div>

    <!-- ── Summary Cards ── -->
    <div class="grid grid-cols-3 gap-4">
      <!-- Income -->
      <div class="flex items-center gap-4 p-5 rounded-2xl bg-[#1a1d2e] border border-white/10 shadow-[0_4px_20px_rgba(0,0,0,0.3)] transition-transform hover:-translate-y-0.5">
        <div class="w-[46px] h-[46px] rounded-xl bg-emerald-500/[0.18] text-emerald-400 flex items-center justify-center flex-shrink-0">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" width="18" height="18">
            <polyline points="23 6 13.5 15.5 8.5 10.5 1 18"/><polyline points="17 6 23 6 23 12"/>
          </svg>
        </div>
        <div>
          <p class="text-xs text-slate-500 mb-1.5">รายรับเดือนนี้</p>
          <p class="text-[22px] font-bold text-emerald-400">฿{{ formatMoney(monthSummary.income) }}</p>
        </div>
      </div>

      <!-- Expense -->
      <div class="flex items-center gap-4 p-5 rounded-2xl bg-[#1a1d2e] border border-white/10 shadow-[0_4px_20px_rgba(0,0,0,0.3)] transition-transform hover:-translate-y-0.5">
        <div class="w-[46px] h-[46px] rounded-xl bg-rose-500/[0.18] text-rose-400 flex items-center justify-center flex-shrink-0">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" width="18" height="18">
            <polyline points="23 18 13.5 8.5 8.5 13.5 1 6"/><polyline points="17 18 23 18 23 12"/>
          </svg>
        </div>
        <div>
          <p class="text-xs text-slate-500 mb-1.5">รายจ่ายเดือนนี้</p>
          <p class="text-[22px] font-bold text-rose-400">฿{{ formatMoney(monthSummary.expense) }}</p>
        </div>
      </div>

      <!-- Net -->
      <div class="flex items-center gap-4 p-5 rounded-2xl bg-[#1a1d2e] border border-white/10 shadow-[0_4px_20px_rgba(0,0,0,0.3)] transition-transform hover:-translate-y-0.5">
        <div class="w-[46px] h-[46px] rounded-xl bg-violet-500/[0.18] text-violet-400 flex items-center justify-center flex-shrink-0">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" width="18" height="18">
            <line x1="12" y1="1" x2="12" y2="23"/><path d="M17 5H9.5a3.5 3.5 0 0 0 0 7h5a3.5 3.5 0 0 1 0 7H6"/>
          </svg>
        </div>
        <div>
          <p class="text-xs text-slate-500 mb-1.5">คงเหลือสุทธิ</p>
          <p class="text-[22px] font-bold" :class="monthSummary.net >= 0 ? 'text-emerald-400' : 'text-rose-400'">
            {{ monthSummary.net >= 0 ? '+' : '' }}฿{{ formatMoney(monthSummary.net) }}
          </p>
        </div>
      </div>
    </div>

    <!-- ── Calendar + Detail ── -->
    <div class="grid gap-5 flex-1" style="grid-template-columns: 1fr 290px;">

      <!-- Calendar Card -->
      <div class="bg-[#1a1d2e] border border-white/10 rounded-[18px] p-[22px] shadow-[0_4px_20px_rgba(0,0,0,0.25)] flex flex-col gap-3.5">
        <!-- Month Nav -->
        <div class="flex items-center justify-between">
          <button
            class="w-8 h-8 rounded-lg border border-white/10 bg-white/[0.05] text-slate-400 flex items-center justify-center cursor-pointer transition-all hover:bg-violet-500/20 hover:text-violet-300 hover:border-violet-500/35"
            @click="prevMonth"
          >
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" width="16" height="16"><polyline points="15 18 9 12 15 6"/></svg>
          </button>
          <h2 class="text-[17px] font-bold text-slate-100">{{ monthLabel }}</h2>
          <button
            class="w-8 h-8 rounded-lg border border-white/10 bg-white/[0.05] text-slate-400 flex items-center justify-center cursor-pointer transition-all hover:bg-violet-500/20 hover:text-violet-300 hover:border-violet-500/35"
            @click="nextMonth"
          >
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" width="16" height="16"><polyline points="9 18 15 12 9 6"/></svg>
          </button>
        </div>

        <!-- Day Headers -->
        <div class="grid grid-cols-7">
          <div
            v-for="(d, i) in thaiDays"
            :key="d"
            :class="['text-center text-[11px] font-semibold py-1.5 uppercase tracking-wide', i === 0 || i === 6 ? 'text-violet-400' : 'text-slate-600']"
          >{{ d }}</div>
        </div>

        <!-- Grid -->
        <div class="grid grid-cols-7 gap-1">
          <div
            v-for="(day, idx) in calendarDays"
            :key="idx"
            :class="[
              'min-h-[78px] rounded-[10px] border p-2 pb-1.5 flex flex-col items-start gap-0.5 transition-all',
              day === null
                ? 'border-transparent cursor-default pointer-events-none'
                : isToday(day) && selectedDate === day
                  ? 'border-violet-500/40 bg-violet-500/[0.14] cursor-pointer'
                  : selectedDate === day
                    ? 'border-violet-500/40 bg-violet-500/[0.14] cursor-pointer'
                    : 'border-transparent cursor-pointer hover:bg-white/[0.05] hover:border-white/[0.09]',
            ]"
            @click="day !== null && (selectedDate = day)"
          >
            <template v-if="day !== null">
              <span
                :class="[
                  'text-[13px] font-medium w-[26px] h-[26px] flex items-center justify-center leading-none flex-shrink-0',
                  isToday(day)
                    ? 'bg-gradient-to-br from-[#6c63ff] to-[#8b84ff] text-white font-bold rounded-full shadow-[0_2px_8px_rgba(108,99,255,0.55)]'
                    : 'text-slate-300',
                ]"
              >{{ day }}</span>
              <div class="flex gap-0.5">
                <span v-if="hasIncome(day)" class="w-[5px] h-[5px] rounded-full bg-emerald-400" />
                <span v-if="hasExpense(day)" class="w-[5px] h-[5px] rounded-full bg-rose-400" />
              </div>
              <div class="flex flex-col gap-px mt-auto">
                <span v-if="hasIncome(day)" class="text-[9px] font-semibold leading-none text-emerald-400">
                  +{{ formatMoney(getEventsForDay(day).filter(e => e.type === 'income').reduce((s,e)=>s+e.amount,0)) }}
                </span>
                <span v-if="hasExpense(day)" class="text-[9px] font-semibold leading-none text-rose-400">
                  -{{ formatMoney(getEventsForDay(day).filter(e => e.type === 'expense').reduce((s,e)=>s+e.amount,0)) }}
                </span>
              </div>
            </template>
          </div>
        </div>
      </div>

      <!-- Detail Panel -->
      <div class="bg-[#1a1d2e] border border-white/10 rounded-[18px] p-[22px] shadow-[0_4px_20px_rgba(0,0,0,0.25)] flex flex-col gap-4 overflow-y-auto">
        <template v-if="selectedDate != null">
          <div class="flex items-start justify-between pb-4 border-b border-white/[0.07]">
            <div class="flex items-baseline gap-2">
              <span class="text-[40px] font-extrabold text-white leading-none">{{ selectedDate }}</span>
              <span class="text-sm text-slate-500">{{ thaiMonths[currentMonth] }}</span>
            </div>
            <div class="text-right flex flex-col gap-0.5">
              <span v-if="selectedDayTotal.income" class="text-sm font-semibold text-emerald-400">+฿{{ formatMoney(selectedDayTotal.income) }}</span>
              <span v-if="selectedDayTotal.expense" class="text-sm font-semibold text-rose-400">-฿{{ formatMoney(selectedDayTotal.expense) }}</span>
            </div>
          </div>

          <div v-if="selectedDayEvents.length" class="flex flex-col gap-2">
            <div
              v-for="(ev, i) in selectedDayEvents"
              :key="i"
              class="flex items-center justify-between px-3.5 py-3 rounded-xl bg-white/[0.04] border border-white/[0.06] transition-colors hover:bg-white/[0.08]"
            >
              <div class="flex items-center gap-2.5">
                <span
                  :class="[
                    'w-2 h-2 rounded-full flex-shrink-0',
                    ev.type === 'income'
                      ? 'bg-emerald-400 shadow-[0_0_6px_rgba(16,185,129,0.6)]'
                      : 'bg-rose-400 shadow-[0_0_6px_rgba(244,63,94,0.6)]',
                  ]"
                />
                <div>
                  <p class="text-[13px] font-medium text-slate-200">{{ ev.title }}</p>
                  <p class="text-[11px] text-slate-500 mt-0.5">{{ ev.category }}</p>
                </div>
              </div>
              <span :class="['text-[13px] font-semibold', ev.type === 'income' ? 'text-emerald-400' : 'text-rose-400']">
                {{ ev.type === 'income' ? '+' : '-' }}฿{{ formatMoney(ev.amount) }}
              </span>
            </div>
          </div>

          <div v-else class="flex-1 flex flex-col items-center justify-center gap-2.5 text-slate-500 text-[13px] text-center py-8">
            <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" width="38" height="38" class="opacity-25">
              <rect x="3" y="4" width="18" height="18" rx="2"/><line x1="16" y1="2" x2="16" y2="6"/><line x1="8" y1="2" x2="8" y2="6"/><line x1="3" y1="10" x2="21" y2="10"/>
            </svg>
            <p>ไม่มีรายการในวันนี้</p>
            <button
              class="px-4 py-1.5 rounded-lg border border-violet-500/40 bg-violet-500/[0.12] text-violet-300 text-xs cursor-pointer transition-colors hover:bg-violet-500/[0.22]"
              @click="showAddModal = true"
            >+ เพิ่มรายการ</button>
          </div>
        </template>

        <div v-else class="flex-1 flex flex-col items-center justify-center gap-3 text-slate-500 text-[13px] text-center">
          <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="1.5" width="48" height="48" class="opacity-15">
            <rect x="3" y="4" width="18" height="18" rx="2"/><line x1="16" y1="2" x2="16" y2="6"/><line x1="8" y1="2" x2="8" y2="6"/><line x1="3" y1="10" x2="21" y2="10"/>
          </svg>
          <p>เลือกวันเพื่อดูรายละเอียด</p>
        </div>
      </div>
    </div>

    <!-- ── Add Event Modal ── -->
    <Transition name="modal">
      <div
        v-if="showAddModal"
        class="fixed inset-0 bg-black/70 backdrop-blur-md flex items-center justify-center z-[200]"
        @click.self="showAddModal = false"
      >
        <div class="bg-[#1a1d2e] border border-white/10 rounded-[20px] p-7 w-[390px] shadow-[0_30px_70px_rgba(0,0,0,0.6)]">
          <div class="flex justify-between items-center mb-[22px]">
            <h3 class="text-[18px] font-bold text-white">เพิ่มรายการ</h3>
            <button
              class="bg-transparent border-none cursor-pointer text-slate-500 p-1 rounded-md flex transition-all hover:text-white hover:bg-white/[0.08]"
              @click="showAddModal = false"
            >
              <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none" stroke="currentColor" stroke-width="2" width="18" height="18">
                <line x1="18" y1="6" x2="6" y2="18"/><line x1="6" y1="6" x2="18" y2="18"/>
              </svg>
            </button>
          </div>

          <!-- Type Tabs -->
          <div class="flex gap-1.5 bg-white/[0.05] p-1 rounded-[10px] mb-5">
            <button
              :class="['flex-1 py-2 rounded-[7px] border-none text-sm font-medium cursor-pointer transition-all', newEvent.type === 'expense' ? 'bg-white/10 text-slate-100' : 'bg-transparent text-slate-500']"
              @click="newEvent.type = 'expense'"
            >รายจ่าย</button>
            <button
              :class="['flex-1 py-2 rounded-[7px] border-none text-sm font-medium cursor-pointer transition-all', newEvent.type === 'income' ? 'bg-white/10 text-slate-100' : 'bg-transparent text-slate-500']"
              @click="newEvent.type = 'income'"
            >รายรับ</button>
          </div>

          <!-- Fields -->
          <div class="mb-3.5">
            <label class="block text-xs text-slate-400 font-medium mb-1.5">รายการ</label>
            <input
              v-model="newEvent.title"
              placeholder="เช่น ค่าอาหาร, เงินเดือน..."
              class="w-full px-3.5 py-2.5 rounded-[10px] border border-white/10 bg-white/[0.05] text-slate-200 text-sm outline-none transition-all placeholder:text-slate-600 focus:border-violet-500/50 focus:bg-violet-500/[0.06]"
            />
          </div>
          <div class="mb-3.5">
            <label class="block text-xs text-slate-400 font-medium mb-1.5">จำนวนเงิน (บาท)</label>
            <input
              v-model="newEvent.amount"
              type="number"
              placeholder="0"
              class="w-full px-3.5 py-2.5 rounded-[10px] border border-white/10 bg-white/[0.05] text-slate-200 text-sm outline-none transition-all placeholder:text-slate-600 focus:border-violet-500/50 focus:bg-violet-500/[0.06]"
            />
          </div>
          <div class="mb-3.5">
            <label class="block text-xs text-slate-400 font-medium mb-1.5">หมวดหมู่</label>
            <input
              v-model="newEvent.category"
              placeholder="เช่น อาหาร, เดินทาง..."
              class="w-full px-3.5 py-2.5 rounded-[10px] border border-white/10 bg-white/[0.05] text-slate-200 text-sm outline-none transition-all placeholder:text-slate-600 focus:border-violet-500/50 focus:bg-violet-500/[0.06]"
            />
          </div>

          <!-- Footer -->
          <div class="flex gap-2.5 mt-[22px]">
            <button
              class="flex-1 py-2.5 rounded-[10px] border border-white/10 bg-transparent text-slate-400 text-sm cursor-pointer transition-all hover:bg-white/[0.06] hover:text-white"
              @click="showAddModal = false"
            >ยกเลิก</button>
            <button
              :class="[
                'flex-1 py-2.5 rounded-[10px] border-none text-sm font-semibold text-white cursor-pointer transition-transform hover:-translate-y-px',
                newEvent.type === 'expense'
                  ? 'bg-gradient-to-br from-rose-500 to-rose-400 shadow-[0_4px_14px_rgba(244,63,94,0.4)]'
                  : 'bg-gradient-to-br from-emerald-500 to-emerald-400 shadow-[0_4px_14px_rgba(16,185,129,0.4)]',
              ]"
              @click="addEvent"
            >บันทึก</button>
          </div>
        </div>
      </div>
    </Transition>
  </main>
</template>

<style scoped>
.modal-enter-active, .modal-leave-active { transition: opacity 0.25s ease; }
.modal-enter-from, .modal-leave-to { opacity: 0; }
.modal-enter-active :deep(.modal-box), .modal-leave-active :deep(.modal-box) { transition: transform 0.25s cubic-bezier(0.4,0,0.2,1); }
.modal-enter-from :deep(.modal-box), .modal-leave-to :deep(.modal-box) { transform: scale(0.94) translateY(8px); }
</style>