import java.util.ArrayList;
import java.util.Arrays;
import java.util.List;
import java.util.Scanner;
import java.util.stream.Collectors;

class Main {
    public static int min_num;

    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int item_number = Integer.parseInt(in.nextLine());

        List<Integer> max_items = Arrays.stream(in.nextLine().split(" ")).map(Integer::parseInt)
                .collect(Collectors.toList());
        List<ArrayList<Integer>> prices = new ArrayList<>();
        for (int i = 0; i < item_number; i++) {
            List<Integer> item_price = Arrays.stream(in.nextLine().split(" ")).map(Integer::parseInt)
                    .collect(Collectors.toList());
            prices.add(new ArrayList<>(item_price));
        }

        int max_profit = 0;
        for (int i = 0; i < prices.size(); i++) {
            int ans = 0;
            for (int j = 1; j < prices.get(i).size(); j++) {
                ans += Math.max(0, prices.get(i).get(j) - prices.get(i).get(j - 1));
            }
            max_profit += ans * max_items.get(i);
        }

        System.out.println(max_profit);

        in.close();
    }
}